FROM golang:1.25-alpine

RUN apk update && apk add git
RUN go install github.com/air-verse/air@latest

WORKDIR /app

ENTRYPOINT ["air"]
CMD ["--build.cmd", "go build -o bin/api cmd/server/main.go", "--build.entrypoint", "./bin/api"]

