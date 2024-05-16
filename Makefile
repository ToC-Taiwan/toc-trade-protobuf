PYTHON=$(shell which python3)
PIP=$(shell which pip3)

compile-go:
	@rm -rf src/golang && mkdir -p src/golang
	@protoc \
	--go_out=src/golang \
	--go-grpc_out=src/golang \
	--proto_path=protos/v3 \
	./protos/v3/*/*.proto
	@. ./scripts/gomod_update.sh

compile-py: check
	@rm -rf src/python/toc_trade_pb && mkdir -p src/python/toc_trade_pb
	@$(PYTHON) -m grpc_tools.protoc \
	--grpc_python_out=src/python/toc_trade_pb \
	--proto_path=protos/v3 \
	./protos/v3/*/*.proto

	@protoc \
	--python_out=pyi_out:src/python/toc_trade_pb \
	--proto_path=protos/v3 \
	./protos/v3/*/*.proto

	@touch src/python/toc_trade_pb/app/__init__.py
	@touch src/python/toc_trade_pb/forwarder/__init__.py

build-py: check
	@rm -rf dist
	@$(PYTHON) -m build

upload-py: check
	@$(PYTHON) -m twine upload --repository pypi dist/*

clean: ## clear virtual environment
	@rm -rf venv

venv: clean ## create virtual environment
	@$(PYTHON) -m venv venv

check: ## check environment
ifneq ($(PYTHON),$(PWD)/venv/bin/python3)
	$(error "Please run 'make venv' first")
endif
	@echo "Venv python version: $(shell $(PYTHON) --version | awk '{print $$2}')"
	@echo "Python path: $(PYTHON)"
	@$(PIP) install -U --no-warn-script-location --no-cache-dir grpcio-tools protobuf build twine

help: ## display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
