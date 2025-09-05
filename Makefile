run:
	go run ./cmd/api/main.go

test:
	go test ./...

test-verbose:
	go test -v ./...
test-cover:
	go test -cover ./...
test-cover-html:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out -o coverage.html