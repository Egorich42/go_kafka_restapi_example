FROM golang:alpine AS build-stage

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN ls cmd

RUN go mod tidy
RUN go mod download 
RUN go build -o cmd/main cmd/main.go

CMD ["cmd/main"]
