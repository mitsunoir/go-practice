.PHONY: build run test

build:
	docker build . -t go-practice-dev

run:
	docker run --name app -p 1323:1323 -v .:/app --rm go-practice-dev

test-author:
	@res=$$(curl -s -X POST "http://localhost:1323/api/v1/authors" -d '{"name": "Jane Doe"}' -H 'Content-Type: application/json'); \
	echo "post: res=$$res"; \
	id=$$(echo "$$res" | jq -r .id); \
	echo "id=$$id"; \
	res=$$(curl -s -X GET "http://localhost:1323/api/v1/authors/$$id"); \
	echo "get: res=$$res";

test-book:
	@res=$$(curl -s -X POST "http://localhost:1323/api/v1/books" -d '{"title": "First Book"}' -H 'Content-Type: application/json'); \
	echo "post: res=$$res"; \
	id=$$(echo "$$res" | jq -r .id); \
	echo "id=$$id"; \
	res=$$(curl -s -X GET "http://localhost:1323/api/v1/books/$$id"); \
	echo "get: res=$$res";

test:
	go test -v ./...
