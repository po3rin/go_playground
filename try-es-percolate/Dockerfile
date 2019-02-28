FROM golang:1.12

WORKDIR /api_server

COPY . .
ENV GO111MODULE=on

RUN go get github.com/pilu/fresh
CMD ["fresh"]
