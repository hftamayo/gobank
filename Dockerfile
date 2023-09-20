FROM golang:1.19.3
FROM golang:alpine

LABEL maintainer="hftamayo"

#RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base

RUN mkdir /opt/hftamayo
RUN mkdir /opt/hftamayo/gobank
WORKDIR /opt/hftamayo/gobank

COPY . ./
COPY .env ./

#name of the executable
RUN go build -o /main

EXPOSE 8001 

#run the exec
CMD ["/main"]
