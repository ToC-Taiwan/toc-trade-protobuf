PYTHON=$(shell which python3)
PIP=$(shell which pip3)
PROTOC_PATH=$(shell which protoc)
PROTOC_INCLUDE_PATH=$(shell dirname $(shell dirname "$(PROTOC_PATH)"))

compile: compile-go compile-py compile-ts compile-dart

compile-go:
	@rm -rf golang && mkdir -p golang
	@$(PROTOC_PATH) \
	--go_out=golang \
	--go-grpc_out=golang \
	--proto_path=protos/v3 \
	./protos/v3/*/*.proto
	@. ./scripts/gomod_update.sh

compile-dart:
	@rm -rf lib && mkdir -p lib
	@dart pub global activate --overwrite protoc_plugin
	@$(PROTOC_PATH) \
	--dart_out=lib \
	--proto_path=protos/v3 \
	./protos/v3/*/*.proto \
	$(PROTOC_INCLUDE_PATH)/include/google/protobuf/empty.proto

compile-py: check
	@rm -rf python/toc_trade_pb && mkdir -p python/toc_trade_pb
	@$(PYTHON) -m grpc_tools.protoc \
	--grpc_python_out=python/toc_trade_pb \
	--python_out=pyi_out:python/toc_trade_pb \
	--proto_path=protos/v3 \
	./protos/v3/*/*.proto

	@./scripts/modify_py_import.sh python/toc_trade_pb
	@touch python/toc_trade_pb/__init__.py
	@touch python/toc_trade_pb/app/__init__.py
	@touch python/toc_trade_pb/forwarder/__init__.py

compile-ts:
	@rm -rf typescript && mkdir -p typescript
	@$(PROTOC_PATH) \
	--proto_path=protos/v3 \
    --ts_opt=no_grpc \
    --ts_opt=no_namespace \
    --ts_out=typescript \
	./protos/v3/*/*.proto

build-py: check
	@rm -rf dist
	@$(PYTHON) -m build

upload-py: check build-py
ifeq ($(PYPI_TOKEN),)
	$(error "PYPI_TOKEN first")
endif
	@$(PYTHON) -m twine upload --repository pypi -p $(PYPI_TOKEN) dist/*

clean:
	@rm -rf venv

venv: clean
	@$(PYTHON) -m venv venv

check: ## check environment
ifneq ($(PYTHON),$(PWD)/venv/bin/python3)
	$(error "Please run 'make venv' first")
endif
	@echo "Venv python version: $(shell $(PYTHON) --version | awk '{print $$2}')"
	@echo "Python path: $(PYTHON)"
	@$(PIP) install -U --no-warn-script-location --no-cache-dir grpcio-tools protobuf build twine
