install:
	go install golang.org/x/tools/cmd/goimports@latest

build:
	go build

test:
	go test ./...

coverage:
	go test -race -coverprofile=coverage.txt -covermode=atomic -coverpkg=./... ./...
	go tool cover -html=coverage.txt

.PHONY: generate
generate: platform importapi ml history

.PHONY: platform
platform:
	java -jar ../rmf-codegen/rmf-codegen.jar generate ../commercetools-api-reference/api-specs/api/api.raml -b platform -o generated -t GO_CLIENT
	goimports -w ./generated/platform/
	go fmt ./generated/platform/
	rm -rf platform/
	mv generated/platform .

.PHONY: importapi
importapi:
	java -jar ../rmf-codegen/rmf-codegen.jar generate ../commercetools-api-reference/api-specs/importapi/api.raml -b importapi -o generated -t GO_CLIENT
	goimports -w ./generated/importapi/
	go fmt ./generated/importapi/
	rm -rf importapi/
	mv generated/importapi .

.PHONY: ml
ml:
	java -jar ../rmf-codegen/rmf-codegen.jar generate ../commercetools-api-reference/api-specs/ml/api.raml -b ml -o generated -t GO_CLIENT
	goimports -w ./generated/ml/
	go fmt ./generated/ml/
	rm -rf ml/
	mv generated/ml .


.PHONY: history
history:
	java -jar ../rmf-codegen/rmf-codegen.jar generate ../commercetools-api-reference/api-specs/history/api.raml -b history -o generated -t GO_CLIENT
	goimports -w ./generated/history/
	go fmt ./generated/history/
	rm -rf history/
	mv generated/history .

.PHONY: connect
connect:
	java -jar ../rmf-codegen/rmf-codegen.jar generate ../commercetools-api-reference/api-specs/connect/api.raml -b connect -o generated -t GO_CLIENT
	goimports -w ./generated/connect/
	go fmt ./generated/connect/
	rm -rf connect/
	mv generated/connect .
