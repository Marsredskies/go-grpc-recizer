FROM golang:1.16

ADD . /go/src/github.com/marsredskies/go-grpc-resizer/server

RUN go install /go/src/github.com/marsredskies/go-grpc-resizer/server

ENTRYPOINT ["/go/bin/server"]

EXPOSE 5300