help: ## display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
.PHONY: help

compile-go:  ### compile-go
	@rm -rf gen/go && mkdir -p gen/go
	@protoc \
	--go_out=gen/go \
	--go_opt=paths=source_relative \
	--go-grpc_out=gen/go \
	--go-grpc_opt=paths=source_relative \
	--proto_path=protos/v3/app \
	--proto_path=protos/v3/forwarder \
	./protos/v3/*/*.proto
.PHONY: compile-go

compile-py:  ### compile-py
	@rm -rf gen/python && mkdir -p gen/python
	@python3 -m grpc_tools.protoc \
	--python_out=gen/python \
	--grpc_python_out=gen/python \
	--proto_path=protos/v3/app \
	--proto_path=protos/v3/forwarder \
	./protos/v3/*/*.proto
.PHONY: compile-py
