FROM golang:1.24-bullseye as base

RUN apt update -y && \
    apt upgrade -y && \
    apt autoremove -y

RUN apt install -y protobuf-compiler tree

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
    && export PATH=$PATH:$(go env GOPATH)/bin

RUN date
WORKDIR /work
COPY ./proto proto/

RUN mkdir -p order && protoc --go_out=./order --go-grpc_out=./order proto/order-engine/order-engine.proto
RUN mkdir -p payment && protoc --go_out=./payment --go-grpc_out=./payment proto/payment-engine/payment-engine.proto
RUN mkdir -p shipping && protoc --go_out=./shipping --go-grpc_out=./shipping proto/shipping-engine/shipping-engine.proto