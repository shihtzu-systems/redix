version := $(shell cat $$HOME/git/redix/version)

build: fmt vet test

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./...

release:
	git add --all
	git commit
	git tag v$(version)

stamp:
	printf "$(version)" > version