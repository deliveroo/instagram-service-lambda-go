FROM golang:1.11.1-alpine

RUN apk add --no-cache alpine-sdk zip

ARG DIR=${GOPATH}/src/github.com/deliveroo/instagram-service-lambda-go
ADD . $DIR
WORKDIR $DIR

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download
