coverage:
	go test ./... -race -coverprofile=coverage.txt -covermode=atomic
