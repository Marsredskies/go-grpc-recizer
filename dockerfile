FROM golang:1.16

ADD . https://github.com/marsredskies/go-grpc-resizer/server

RUN go install github.com/marsredskies/go-grpc-resizer/server@latest

ENTRYPOINT ["/go/bin/server"]

EXPOSE 5300
