FROM golang:alpine

RUN apk update && apk add --no-cache gcc musl-dev

WORKDIR /linkaja-test

COPY . .

RUN go build -o linkaja-test

CMD ["/linkaja-test/linkaja-test"]