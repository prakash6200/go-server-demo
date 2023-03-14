FROM golang:1.14

RUN mkdir /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go build 

CMD ["/app/main"]
