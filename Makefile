.PHONY: all clean build deps install uninstall

PREFIX = /usr/local
GOBUILD = go build
GOGET = go get -d -v

all: deps clean build

clean:
	rm -rf bin

build:
	$(GOBUILD) -o bin/teambot main.go

deps:
	$(GOGET) ./...

install: bin/teambot
	install -m0755 bin/teambot $(PREFIX)/bin/teambot

uninstall: $(PREFIX)/bin/teambot
	rm -rf $(PREFIX)/bin/teambot

