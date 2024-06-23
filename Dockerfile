FROM golang:1.22 as builder

RUN mkdir /app
WORKDIR /app

COPY . .

RUN go mod download
RUN go mod verify

ARG APP_VERSION
RUN go build -o bin/avion_server -ldflags "-X main.Version=$APP_VERSION" ./cmd

RUN wget -qO- https://github.com/golang-migrate/migrate/releases/download/v4.16.0/migrate.linux-amd64.tar.gz | tar xvz -C /usr/local/bin/ \
    && chmod +x /usr/local/bin/migrate

COPY migrations /migrations

FROM debian:testing-slim

COPY --from=builder /migrations /migrations
COPY --from=builder /app/bin/avion_server /app/avion_server
COPY --from=builder /usr/local/bin/migrate /usr/local/bin/migrate

RUN mkdir -p /app

WORKDIR /app

CMD ./avion_server
