GOARCH=amd64
BINARY=xkpassgo

build: linux darwin windows

linux:
	CGO_ENABLED=0 GOARCH=${GOARCH} GOOS=linux go build -ldflags="-s -w" -o bin/${BINARY}-linux-${GOARCH} ./cmd/xkpassgo/main.go

darwin:
	CGO_ENABLED=0 GOARCH=${GOARCH} GOOS=darwin go build -ldflags="-s -w" -o bin/${BINARY}-darwin-${GOARCH} ./cmd/xkpassgo/main.go

windows:
	CGO_ENABLED=0 GOARCH=${GOARCH} GOOS=windows go build -ldflags="-s -w" -o bin/${BINARY}-windows-${GOARCH}.exe ./cmd/xkpassgo/main.go

.PHONY: build
