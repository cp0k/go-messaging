FROM golang:1.11.5
WORKDIR demo
RUN go mod init demo
RUN go get github.com/nats-io/go-nats@v1.7.0
COPY main.go .
CMD ["go", "run", "main.go"]
