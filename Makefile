.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/lambda lambda/main.go
	go build -o bin/cli cmd/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
