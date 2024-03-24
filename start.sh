#!/bin/sh

set -e

echo "start server"
source /app/app.env

/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"