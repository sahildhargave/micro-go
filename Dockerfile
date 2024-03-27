# Build stage
FROM golang:1.22.0 AS builder

WORKDIR /app

# Copy only necessary files for building
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go application
RUN go build -o main main.go

# Run stage
FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/main .


COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration

# Set executable permissions for start.sh
RUN chmod +x /app/start.sh


EXPOSE 8080 9090

CMD ["./main"]
