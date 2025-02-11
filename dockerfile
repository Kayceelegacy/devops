FROM golang:1.19 AS build-stage

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o /main

EXPOSE 8080

CMD ["/main"]

