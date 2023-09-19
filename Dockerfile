FROM golang:1.19.3
#FROM golang:alpine

LABEL maintainer="hftamayo"

RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base

RUN mkdir /opt/hftamayo
RUN mkdir /opt/hftamayo/gobank
WORKDIR /opt/hftamayo/gobank

COPY . /opt/hftamayo/gobank
COPY .env /opt/hftamayo/gobank

RUN go build -o main .

CMD["./main"]
