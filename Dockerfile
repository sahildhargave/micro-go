
# Build stage
FROM golang:1.21-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

EXPOSE 8081

CMD ["/app/main"]
