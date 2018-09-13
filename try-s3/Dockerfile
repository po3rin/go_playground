FROM golang:1.11.0

WORKDIR /go/src/try-s3

COPY . .

RUN mkdir -p ~/.aws
RUN mv ./.aws/credentials ~/.aws/

RUN go get -d -v ./...