.DEFAULT_GOAL: all

.PHONY: all
all: generate-skeletons

.PHONY: generate-skeletons
generate-skeletons:
	docker build -t proto-generator:latest .
	docker run -v ./:/artifacts -it proto-generator:latest /bin/bash -c "cp -rf github.com/sklrsn/gRPC-kube/* /artifacts/"