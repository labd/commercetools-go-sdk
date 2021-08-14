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
	go fmt ./generated/platform/
	rm -rf platform/
	mv generated/platform .

.PHONY: importapi
importapi:
	java -jar ../rmf-codegen/rmf-codegen.jar generate ../commercetools-api-reference/api-specs/importapi/api.raml -b importapi -o generated -t GO_CLIENT
	go fmt ./generated/importapi/
	rm -rf importapi/
	mv generated/importapi .

.PHONY: ml
ml:
	java -jar ../rmf-codegen/rmf-codegen.jar generate ../commercetools-api-reference/api-specs/ml/api.raml -b ml -o generated -t GO_CLIENT
	go fmt ./generated/ml/
	rm -rf ml/
	mv generated/ml .


.PHONY: history
history:
	java -jar ../rmf-codegen/rmf-codegen.jar generate ../commercetools-api-reference/api-specs/history/api.raml -b history -o generated -t GO_CLIENT
	go fmt ./generated/history/
	rm -rf history/
	mv generated/history .
