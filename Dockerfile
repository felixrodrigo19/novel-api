FROM golang:1.21.11-alpine3.20

LABEL authors="Rodrigo dos Santos Felix"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

COPY controllers/ ./controllers

COPY database/ ./database

COPY models/ ./models

COPY routes/ ./routes

RUN CGO_ENABLED=0 GOOS=linux go build -o /novel-api

EXPOSE 8000

CMD ["/novel-api"]

