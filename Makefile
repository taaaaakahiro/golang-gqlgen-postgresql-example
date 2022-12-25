run:
	go run ./cmd/api/main.go

validate:
	go run ./cmd/gqlgenerate/main.go

test:
	go test ./...

generate:
	go generate ./...