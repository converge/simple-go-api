# todo: add stage for building the binary
#       run with less privileges
FROM golang:1.19-alpine3.16

WORKDIR /opt/app

COPY . .

RUN cd cmd/http && go build -o ../../simple-go-api

CMD ["./simple-go-api"]
