# build stage
FROM golang:1.13 AS build-env

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
# it will take the flags from the environment
RUN go build -o warehouse_api

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /warehouse_api /app/
ENTRYPOINT ./warehouse_api