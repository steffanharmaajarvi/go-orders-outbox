#!/bin/sh

set -e

echo "⏳ Waiting for PostgreSQL to be ready..."

until pg_isready -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER"; do
  sleep 1
done

echo "✅ PostgreSQL is ready. Running migrations..."

migrate -path /app/migrations -database "$DB_DSN" up

echo "✅ Migrations applied successfully."
