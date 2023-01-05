test:
	go test ./...

build:
	go build .

test-cov:
	go test ./... -coverprofile=coverage.out
