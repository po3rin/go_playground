FROM golang:1.12 as builder
WORKDIR /go/api
COPY . .
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest as prod
EXPOSE 8080
RUN apk --no-cache add ca-certificates
WORKDIR /api
COPY --from=builder /go/api/ .
RUN pwd
CMD ["./try-docker-target"]


FROM golang:1.12 as dev
EXPOSE 8080
WORKDIR /go/api
ENV GO111MODULE=on
COPY . .
RUN go get github.com/pilu/fresh
CMD ["fresh"]
