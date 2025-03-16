.DEFAULT_GOAL: all

.PHONY: all
all: generate-skeletons

.PHONY: generate-skeletons
generate-skeletons: clean
	docker build -t proto-generator:latest .
	docker run -v ./:/artifacts -it proto-generator:latest /bin/bash -c "cp -rf order/github.com/sklrsn/gRPC-defs/order /artifacts/ && \
	cp -rf payment/github.com/sklrsn/gRPC-defs/payment /artifacts/ &&\
	cp -rf shipping/github.com/sklrsn/gRPC-defs/shipping /artifacts/"

clean:
	rm -rf ./artifacts
	rm -rf ./order
	rm -rf ./payment
	rm -rf ./shipping