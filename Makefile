PROJECT:=micro-service-gen-tool

.PHONY: build

build-linux, build:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o mss-boot-generator main.go
build-windows:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -o mss-boot-generator.exe main.go
build-darwin:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -o mss-boot-generator main.go
