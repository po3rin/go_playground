FROM golang:1.10.0 as builder

WORKDIR /go/src/go-playground/try-logrus-stackdriver

COPY . .

RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /go/src/go-playground/try-logrus-stackdriver .

CMD ["./try-logrus-stackdriver"]
