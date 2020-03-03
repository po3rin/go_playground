try Go + RabbitMQ

## QuickStart

run rabbitmq container.

```bash
$ docker-compose up -d
```

run publisher.

```bash
$ go run send.go
```

run consumer.

```bash
$ go run receive.go
```
