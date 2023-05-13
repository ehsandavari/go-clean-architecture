help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"; printf "\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
.PHONY: help

run: linter-golangci swagger ### run
	go mod tidy && go mod download && \
	go run .
.PHONY: run

swagger: ### init swagger
	swag fmt
	swag init --parseDependency --parseInternal --parseDepth 1 --generalInfo ./presentation/api/v1/v1.go --output ./presentation/api/docs --instanceName v1 --generatedTime true
.PHONY: swagger

linter-golangci: ### check by golangci linter
	golangci-lint run ./...
.PHONY: linter-golangci

test: ### run unit-test and integration-test
	go test -v -cover -race ./... && \
    go clean -testcache && \
    go test -v ./test/...
.PHONY: test

install-deps: ### install all dependencies
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/golang/mock/mockgen@latest
	go install golang.org/x/tools/cmd/stringer@latest
.PHONY: install-deps


proto: ### init proto
	protoc --go_out=presentation/grpc/proto/bankAccount/ --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=presentation/grpc/proto/bankAccount/ presentation/grpc/proto/bankAccount/*.proto
.PHONY: proto
