test:
	@go test -v ./...

run:
	@go run main.go

build:
	@go build -o github.com/oniharnantyo/golang-backend-example main.go

docker:
	@docker build -t oniharnantyo/github.com/oniharnantyo/golang-backend-example .

docker-compose:
	@docker-compose up -d

docker-push:
	@docker push oniharnantyo/github.com/oniharnantyo/golang-backend-example