FROM golang:1.17-alpine3.14

RUN apk update && apk add curl librdkafka gcc musl-dev

CMD ["go", "run", "-tags", "musl", "main.go", "serve", "--debug"]
