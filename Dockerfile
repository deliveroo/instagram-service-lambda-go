FROM goland:1.11.1-alphine

RUN apk add --no-cache alpine-sdk zip

ARG DIR=${GOPATH}/src/github.com/deliveroo/instagram-service-lambda-go
ADD . $DIR
WORKDIR $DIR