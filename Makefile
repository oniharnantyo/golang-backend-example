test:
	@go test -v ./...

run:
	@go run main.go

build:
	@go build -o golang-backend-example main.go

docker:
	@docker build -t oniharnantyo/golang-backend-example .

docker-compose:
	@docker-compose up -d

docker-push:
	@docker push oniharnantyo/golang-backend-example