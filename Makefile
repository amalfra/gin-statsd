.PHONY: all

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test -v ./... -coverprofile=profile.cov

build: fmt vet test
