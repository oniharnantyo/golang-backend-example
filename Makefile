test:
	@go test -v ./...

run:
	@go run main.go

build:
	@go build -o linkaja-test main.go

docker:
	@docker build -t oniharnantyo/linkaja-test .

docker-compose:
	@docker-compose up -d

docker-push:
	@docker push oniharnantyo/linkaja-test