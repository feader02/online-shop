FROM golang:1.22 as builder

RUN mkdir app
WORKDIR /app

COPY . .

RUN go mod download
RUN go mod verify

ARG APP_VERSION
RUN go build -o bin/avion_server -ldflags "-X main.Version=$APP_VERSION" ./cmd


FROM busybox

RUN mkdir app

COPY --from=builder /app/bin/* /app
COPY migrations /migrations

CMD ["./app/avion_server"]