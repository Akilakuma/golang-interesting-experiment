FROM golang:1.11-alpine

COPY ./main.go /go/src/rpctest/main.go
WORKDIR /go/src/rpctest
RUN go build -o trpc .

ENTRYPOINT [ "./trpc"]
