# syntax=docker/dockerfile:1

FROM golang:1.21-alpine

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY api ./api
RUN go build

EXPOSE 8080

CMD [ "/app/healthcheck" ]
