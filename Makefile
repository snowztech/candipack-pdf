.PHONY: fmt lint run build clean

APP_NAME=candipack-pdf

fmt:
	gofmt -s -l -w .

lint: fmt
	golangci-lint run

run:
	go run *.go server

build:
	go build -o bin/$(APP_NAME)

clean:
	rm -rf output
	rm -rf bin/
	@echo "cleaned!"