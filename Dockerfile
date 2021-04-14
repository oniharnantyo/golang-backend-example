FROM golang:alpine

RUN apk update && apk add --no-cache gcc musl-dev

WORKDIR /golang-backend-example

COPY . .

RUN go build -o golang-backend-example

CMD ["/golang-backend-example/golang-backend-example"]