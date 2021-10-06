# syntax=docker/dockerfile:1
FROM golang:1.17

RUN go get github.com/segmentio/kafka-go
RUN go get github.com/joho/godotenv
RUN go get go.mongodb.org/mongo-driver/bson
RUN go get go.mongodb.org/mongo-driver/mongo
RUN go get go.mongodb.org/mongo-driver/mongo/options


COPY . .
WORKDIR ./src/consumer-mod/consumer-pkg
RUN go build .

# RUN go mod tidy
# RUN go install -v ./...

# CMD ["app"]
