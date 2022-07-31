FROM golang:1.18-alpine3.16

WORKDIR /opt/app

COPY . .

RUN go build -o simple-go-api

CMD ["./simple-go-api"]