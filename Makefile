GOBUILD = go build
GOGET = go get -d -v
clean:
	rm -rf bin

build:
	$(GOBUILD) -o bin/teambot main.go

deps:
	$(GOGET) ./...

