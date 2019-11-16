# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ADD . .
RUN go build -o warehouse_api

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /warehouse_api /app/
ENTRYPOINT ./warehouse_api