test:
	go test ./...

coverage:
	go test -race -coverprofile=coverage.txt -covermode=atomic -coverpkg=./commercetools ./commercetools
	go tool cover -html=coverage.txt

commercetools-api-reference:
	@git clone https://github.com/commercetools/commercetools-api-reference.git ../commercetools-api-reference

code-generate/prepared.raml: commmercetools-api-reference
	python code-generate/prepare_raml ../commercetools-api-reference/api.raml ./code-generate/prepared.raml

generate: code-generate/prepared.raml
	go run code-generate/*.go ./code-generate/prepared.raml
	goimports -w commercetools/types_*.go
