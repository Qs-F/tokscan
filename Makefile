test:
	go test ./...

build:
	go build .

cov:
	go test ./... -coverprofile=coverage.out
