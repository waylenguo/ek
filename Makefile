APP=ek
GOOS=linux
Version=v1.0

.PHONY: all build run gotool clean help

build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=amd64 go build -ldflags "-X 'github.com/ek/pkg/version.Version=${Version}'"  -o ${APP} main.go

clean:
	go clean