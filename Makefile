.PHONY: run
run: swag
	go mod tidy && go mod download && \
	go run ./cmd/main.go

.PHONY: swag
swag:
	swag fmt && swag init -g internal/delivery/restapi/server.go \
        -o internal/docs

.PHONY: test
test:
	go test -v -coverprofile cover.out ./...
	go tool cover -html cover.out -o cover.html
	rm cover.out