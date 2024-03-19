# Build stage
FROM golang:1.21-alpine3.18 AS builder

WORKDIR /app

# Copy only necessary files for building
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main main.go



# Download and install Golang Migrate
RUN apk add --no-cache curl && \
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.0/migrate.linux-amd64.tar.gz -o migrate.tar.gz && \
	tar -xf migrate.tar.gz && \
	mv migrate.linux-amd64 /usr/local/bin/migrate && \
	rm migrate.tar.gz

# Run stage
FROM alpine:3.18

# Set DNS resolver
WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /usr/local/bin/migrate /usr/local/bin/migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration

EXPOSE 8080 9090

# Command to run the Go application
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]