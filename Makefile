GO=go
all: fmt vet build

test: vet
	@echo "test"
	@go test -v ./...

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

vet:
	@echo "vet"
	@go vet ./...