FROM golang:latest 
ADD . /go/src
WORKDIR /go/src
RUN go get github.com/nats-io/go-nats
RUN go build -o tentacle .
EXPOSE 8080
ENTRYPOINT ["/go/src/tentacle"]