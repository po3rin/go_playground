FROM golang:1.10.0 as builder

WORKDIR /go/src/prism/api-server

COPY . .
COPY ./_outpkg/ /go/src/prism/

RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /go/src/prism/api-server .

CMD ["./api-server"]