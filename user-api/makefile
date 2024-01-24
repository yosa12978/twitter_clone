build:
	@go build -o bin/user-service ./main.go
	@cp .env bin

run: build
	@./bin/user-service

test:
	@go test