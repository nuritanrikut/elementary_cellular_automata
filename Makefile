BINARY_NAME=automata

build: dependencies
	@go build -o bin/$(BINARY_NAME) -v

dependencies:
	@go mod tidy

run: build
	@./bin/$(BINARY_NAME)

clean:
	@rm -rf bin/$(BINARY_NAME)
