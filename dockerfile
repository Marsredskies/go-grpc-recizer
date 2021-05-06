FROM golang:1.16

ADD . https://github.com/marsredskies/go-grpc-resizer

RUN go install github.com/marsredskies/go-grpc-resizer/server_side@latest

ENTRYPOINT ["/go/bin/server"]

EXPOSE 5300

