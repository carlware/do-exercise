
test:
	go test --race ./...

build:
	go build -o exif *.go
