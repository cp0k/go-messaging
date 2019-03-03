FROM golang:1.12
RUN go get golang.org/x/tools/cmd/present

# code example dependencies
RUN go get github.com/nats-io/go-nats

# telnet to play with text based protocols
RUN apt update
RUN apt install -y telnet

# present files and assets
WORKDIR present
COPY go-messaging.slide .
COPY visual visual/
COPY code code/

# start present's HTTP server
EXPOSE 3999
CMD ["present", "-notes", "-http", "0.0.0.0:3999", "-orighost", "localhost"]
