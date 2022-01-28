FROM golang:onbuild

WORKDIR /app

ADD . /app

RUN go build -o tcp-scannet
