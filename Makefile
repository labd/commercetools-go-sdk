PWD := $(shell pwd)
GOPATH := $(shell go env GOPATH)

test:
	go test ./...

coverage:
	go test -race -coverprofile=coverage.txt -covermode=atomic -coverpkg=./... ./...
	go tool cover -html=coverage.txt

getdeps:
	@echo "Installing golint" && go get -u github.com/golang/lint/golint
	@echo "Installing gocyclo" && go get -u github.com/fzipp/gocyclo
	@echo "Installing deadcode" && go get -u github.com/remyoudompheng/go-misc/deadcode
	@echo "Installing misspell" && go get -u github.com/client9/misspell/cmd/misspell
	@echo "Installing ineffassign" && go get -u github.com/gordonklaus/ineffassign

verifiers: getdeps vet fmt lint cyclo deadcode spelling

vet:
	@echo "Running $@"
	@go tool vet -atomic -bool -copylocks -nilfunc -printf -shadow -rangeloops -unreachable -unsafeptr -unusedresult commercetools
	@go tool vet -atomic -bool -copylocks -nilfunc -printf -shadow -rangeloops -unreachable -unsafeptr -unusedresult service

fmt:
	@echo "Running $@"
	@gofmt -d commercetools
	@gofmt -d service

lint:
	@echo "Running $@"
	@${GOPATH}/bin/golint -set_exit_status github.com/labd/commercetools-go-sdk/commercetools...
	@${GOPATH}/bin/golint -set_exit_status github.com/labd/commercetools-go-sdk/service...

ineffassign:
	@echo "Running $@"
	@${GOPATH}/bin/ineffassign .

cyclo:
	@echo "Running $@"
	@${GOPATH}/bin/gocyclo -over 200 commercetools
	@${GOPATH}/bin/gocyclo -over 200 service

deadcode:
	@echo "Running $@"
	@${GOPATH}/bin/deadcode -test $(shell go list ./...) || true

spelling:
	@${GOPATH}/bin/misspell -locale US -error `find commercetools/`
	@${GOPATH}/bin/misspell -locale US -error `find service/`
	@${GOPATH}/bin/misspell -locale US -error `find testutil/`
