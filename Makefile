test:
	go test ./...

coverage:
	go test -race -coverprofile=coverage.txt -covermode=atomic -coverpkg=./commercetools ./commercetools
	go tool cover -html=coverage.txt

code-generate/prepared.raml:
	python ./code-generate/prepare_yaml.py ../commercetools-api-reference/api.raml > ./code-generate/prepared.raml
	# Incorrect type string, need to properly set FieldType type
	yq w -i code-generate/prepared.raml types.FieldDefinition.properties.type.type FieldType

generate: code-generate/prepared.raml
	go run code-generate/*.go ./code-generate/prepared.raml
	goimports -w commercetools/types_*.go
