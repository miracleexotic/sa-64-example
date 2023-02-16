# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

WORKDIR /backend

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build -o /sa-64-backend

EXPOSE 8080

CMD [ "/sa-64-backend" ]


