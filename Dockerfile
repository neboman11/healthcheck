# syntax=docker/dockerfile:1

FROM golang:1.26-alpine3.22

WORKDIR /app
COPY gom.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY api ./api
RUN go build

EXPOSE 8080

CMD [ "/app/healthcheck" ]
