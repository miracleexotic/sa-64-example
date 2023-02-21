# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

RUN apk --update add \
    go \
    musl-dev

WORKDIR /backend

RUN apk --update add \
    util-linux-dev

COPY backend/go.mod ./
COPY backend/go.sum ./
RUN go mod download

COPY backend/. .
RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o /sa-64-backend

EXPOSE 8080

CMD [ "/sa-64-backend" ]


