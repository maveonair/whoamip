.PHONY: clean build dev test

VERSION=0.5.0

clean:
	rm -rf ./dist/*

build: clean
	CGO_ENABLED=0  go build -o ./dist/whoamip -a -ldflags '-s' -installsuffix cgo cmd/whoamip/main.go

dev:
	go run cmd/whoamip/main.go

test:
	go test ./...

release:
	docker buildx build --platform linux/amd64,linux/arm64 -t maveonair/whoamip:$(VERSION) -t maveonair/whoamip:latest --push .
