package main

import (
	"context"
	"strconv"
	"warehouse/app/api/handlers"
	"warehouse/app/api/logging"
	"warehouse/app/application/bus"
	"warehouse/app/application/usecase"
	"warehouse/app/common"
	"warehouse/app/domain/service"
	ibus "warehouse/app/infrastructure/bus"
	"warehouse/app/infrastructure/connection"
	"warehouse/app/infrastructure/rabbitmq"
	"warehouse/app/infrastructure/repository"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func parseInt(idsStr []string) []int {
	res := make([]int, len(idsStr))
	for i, v := range idsStr {
		id, _ := strconv.Atoi(v)
		res[i] = id
	}
	return res
}

func connect(connection string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(connection)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func getAppPrefix() string {
	path := common.GetOsEnvOrDefault("PATH_BASE", "/")
	return path
}

func NewLogger() *log.Logger {
	log.SetFormatter(&log.JSONFormatter{})
	var logger = log.New()
	logger.Formatter = &log.JSONFormatter{}
	return logger
}

func main() {
	ctx := context.Background()
	log := NewLogger()
	r := gin.New()
	r.Use(logging.Logger(log), gin.Recovery())
	router := r.Group(getAppPrefix())
	rClient, err := rabbitmq.NewRabbitMqClient(common.GetOsEnvOrDefault("RABBITMQ_CONNECTION", "amqp://guest:guest@rabbitmq:5672/"))
	if err != nil {
		log.WithError(err).Fatal("Cannot connect to rabbitmq")
	}
	defer rClient.Close()
	stateRepository := repository.NewWarehouseStateRepository(connection.NewRedisClient(ctx))
	stateService := service.NewWarehouseStateService(stateRepository)
	useCase := usecase.NewWarehouseStateUseCaseUseCase(stateRepository, stateService)
	subscriber := ibus.NewRabbitMqMessageSubscriber(rClient, common.GetOsEnvOrDefault("RABBITMQ_EXCHANGE", "catalog"), common.GetOsEnvOrDefault("RABBITMQ_QUEUE", "warehouse"), common.GetOsEnvOrDefault("RABBITMQ_TOPIC", "products"))
	api := handlers.NewWarehouseApi(router, useCase)
	go bus.HandleProductCreated(ctx, subscriber, stateRepository)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	api.Start(ctx)
	_ = r.Run() // listen and serve on 0.0.0.0:8080
}
