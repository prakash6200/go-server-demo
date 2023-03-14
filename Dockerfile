FROM golang:1.14.0-apline3.9

RUN mkdir /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go build -o main

CMD ["/app/main"]
