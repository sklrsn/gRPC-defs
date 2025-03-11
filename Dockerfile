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

RUN protoc --go_out=. --go-grpc_out=. proto/order-engine/order-engine.proto
RUN protoc --go_out=. --go-grpc_out=. proto/payment-engine/payment-engine.proto
RUN protoc --go_out=. --go-grpc_out=. proto/shipping-engine/shipping-engine.proto