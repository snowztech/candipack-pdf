.PHONY: fmt lint run build clean

APP_NAME=candipack-pdf

fmt:
	gofmt -s -l -w .

lint: fmt
	golangci-lint run

logs:
	docker-compose logs -f

build:
	docker-compose build

run:
	docker-compose up

clean:
	rm -rf output
	rm -rf bin/
	@echo "cleaned!"