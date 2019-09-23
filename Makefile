GOARCH=amd64
BINARY=xkpassgo

build: linux darwin windows

linux:
	GO111MODULE=on CGO_ENABLED=0 GOARCH=${GOARCH} GOOS=linux packr2 build -ldflags="-s -w" -mod=vendor -o ../../bin/${BINARY}-linux-${GOARCH} ./cmd/xkpassgo/main.go

darwin:
	GO111MODULE=on CGO_ENABLED=0 GOARCH=${GOARCH} GOOS=darwin packr2 build -ldflags="-s -w" -mod=vendor -o ../../bin/${BINARY}-darwin-${GOARCH} ./cmd/xkpassgo/main.go

windows:
	GO111MODULE=on CGO_ENABLED=0 GOARCH=${GOARCH} GOOS=windows packr2 build -ldflags="-s -w" -mod=vendor -o ../../bin/${BINARY}-windows-${GOARCH}.exe ./cmd/xkpassgo/main.go

.PHONY: build