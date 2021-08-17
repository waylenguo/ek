APP=ek
GOOS=linux

.PHONY: all build run gotool clean help

build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=amd64 go build -o ${APP} main.go

clean:
	go clean