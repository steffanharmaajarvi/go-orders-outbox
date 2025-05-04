# Build stage
FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add curl postgresql-client

# Скачиваем бинарник golang-migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz
RUN mv migrate /usr/local/bin/migrate

COPY ./docker/migrator/wait-and-migrate.sh ./
RUN chmod +x ./wait-and-migrate.sh

ENTRYPOINT ["./wait-and-migrate.sh"]
