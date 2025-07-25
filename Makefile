build:
	golangci-lint run
	golangci-lint fmt
	go build ./...
	go test ./...

unit-tests:
	go test -short ./...