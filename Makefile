test:
	go test ./...

coverage:
	go test -race -coverprofile=coverage.txt -covermode=atomic -coverpkg=./... ./...
	go tool cover -html=coverage.txt

generate:
	go run code-generate/*.go
	goimports -w commercetools/types_*.go
