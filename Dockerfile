# syntax=docker/dockerfile:1

FROM golang:1.26-alpine3.22

RUN apk upgrade --no-cache

WORKDIR /app
COPY go.mod ./
