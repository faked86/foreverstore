build:
	@go build -o bin/fst
run: build
	@./bin/fst
test:
	@go test ./... -v
