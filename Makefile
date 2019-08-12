GOBUILD = go build

clean:
	rm -rf bin

build:
	$(GOBUILD) -o bin/teambot main.go
