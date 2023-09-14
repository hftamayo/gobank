#FROM golang:1.19.3
FROM golang:alpine

LABEL maintainer="hftamayo"

RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base

RUN mkdir /opt/hftamayo/gobank
WORKDIR /opt/hftamayo/gobank

COPY . .
COPY .env .

RUN go get -d -v ./..
RUN go install -v ./..
RUN go build -o /build

RUN go run main.go