#!/bin/sh

set -e  # This ensures that the script exits immediately if any command returns a non-zero exit status.

echo "Starting server..."

# Uncomment the following lines if you need to source environment variables from app.env
# source /app/app.env

# Run database migrations (if applicable)
# /app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "Starting the application..."
exec "$@"
