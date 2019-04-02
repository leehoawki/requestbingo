FROM golang:latest
WORKDIR $GOPATH/src/requestbingo
ADD . $GOPATH/src/requestbingo
RUN go get ./...
RUN go build .
EXPOSE 8080
ENTRYPOINT  ["./requestbingo"]