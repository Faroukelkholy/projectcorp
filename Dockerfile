FROM golang:1.15

WORKDIR /go/src/projectcorp

COPY . .

RUN go mod tidy

EXPOSE 3000

CMD ["go","run","main.go"]