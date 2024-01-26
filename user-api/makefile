build:
	@go build -o bin/user-service ./main.go
	@cp .env bin

run: build
	@./bin/user-service

test:
	@go test

lint:
	@golangci-lint run

docker:
	docker build -t yosaa5782/twitter-user-api .
	docker run --rm --name twitter-user-api -p 5000:5000 --network twitter-net -d yosaa5782/twitter-user-api
