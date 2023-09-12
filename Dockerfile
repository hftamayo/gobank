FROM golang:1.19.3

WORKDIR /opt/hftamayo/gobank

RUN go install 

COPY . .
RUN go run main.go
