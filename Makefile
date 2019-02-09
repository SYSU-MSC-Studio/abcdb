GO=go
all: fmt vet build

test: vet
	@echo "test"
	@go test -v ./...
	@go test -v ./vendor/...

build:
	@echo "build"
	@go build

clean:
	@echo "clean"
	@go clean -i

deploy:
	@echo "deploy"
	@go install

fmt:
	@echo "fmt"
	@go fmt ./...
	@go fmt ./vendor/...

vet:
	@echo "vet"
	@go vet ./...
	@go vet ./vendor/...

help:
	@echo "use make to build"
	@echo "use make test to test"