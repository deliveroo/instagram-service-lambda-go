FROM golang:1.11.1-alpine

RUN apk add --no-cache alpine-sdk zip

ENV GO111MODULE=on

ARG DIR=${GOPATH}/src/github.com/deliveroo/instagram-service-lambda-go
ADD . $DIR
WORKDIR $DIR