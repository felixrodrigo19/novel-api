FROM golang:1.21.11-alpine3.20

LABEL authors="Rodrigo dos Santos Felix"

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /novel-api
