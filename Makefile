test:
	go test ./...
	@make build
	@make clean

build:
	go build ./cmd/tokscan

cov:
	go test ./... -coverprofile=coverage.out
	@make clean

clean:
	@rm -rf ./tokscan ./coverage.out
