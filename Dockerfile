FROM golang:1.16.4 AS BUILD

RUN mkdir /app
WORKDIR /app

ADD go.mod .
ADD go.sum .
RUN go mod download

ADD . ./
RUN go test -bench=. ./... -benchmem -benchtime=10s
