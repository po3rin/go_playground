FROM golang:1.11 AS go111

WORKDIR /go/src/performance
COPY . .
ENV GO111MODULE=on

FROM golang:1.12 AS go112

WORKDIR /go/src/performance
COPY . .
ENV GO111MODULE=on
