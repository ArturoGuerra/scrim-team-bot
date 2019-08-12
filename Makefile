.PHONY: all clean build deps

GOBUILD = go build
GOGET = go get -d -v

all: deps clean build

clean:
	rm -rf bin

build:
	$(GOBUILD) -o bin/teambot main.go

deps:
	$(GOGET) ./...
