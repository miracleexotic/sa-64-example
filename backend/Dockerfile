# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

RUN apk --update add \
    go \
    musl-dev

WORKDIR /backend

RUN apk --update add \
    util-linux-dev

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o /sa-64-backend

EXPOSE 8080

CMD [ "/sa-64-backend" ]


