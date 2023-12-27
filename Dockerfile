FROM golang:1.21.5

WORKDIR /go/src/echo

COPY . .
RUN go mod download

RUN go build -o echo .

EXPOSE 8888

CMD ["./echo"]