# This controls the location of the cache.
PROJECT := advertisement

# This controls the version of buf to install and use.
BUF_VERSION := 1.30.1

PROTOC_VERSION := 25.3
PROTOC_GEN_DOC_VERSION := 1.5.1
PROTOC_GEN_GO_VERSION := 1.33.0
PROTOC_GEN_GO_GRPC_VERSION := 1.3.0
PROTOC_GEN_GO_GRPC_GATEWAY_VERSION := 2.19.1
PROTOC_GEN_GRPC_JAVA_VERSION := 1.63.0
PROTOC_GEN_JS_VERSION := 3.21.2
GRPC_TOOLS_VERSION := 1.12.4
TS_PROTOC_GEN_VERSION := 0.15.0
WIRE_VERSION := 0.6.0

UNAME_OS := $(shell uname -s)
UNAME_ARCH := $(shell uname -m)
ifeq ($(UNAME_OS),Darwin)
	PLATFORM := osx
	ifeq ($(UNAME_ARCH),arm64)
		PROTOC_ARCH := aarch_64
	else
		PROTOC_ARCH := x86_64
	endif
else
	PROTOC_ARCH := $(UNAME_ARCH)
endif
ifeq ($(UNAME_OS),Linux)
	PLATFORM := linux
endif

# Buf will be cached to ~/.cache/buf-example.
CACHE_BASE := $(HOME)/.cache/$(PROJECT)
# This allows switching between i.e a Docker container and your local setup without overwriting.
CACHE := $(CACHE_BASE)/$(UNAME_OS)/$(UNAME_ARCH)
# The location where buf will be installed.
CACHE_BIN := $(CACHE)/bin
# Marker files are put into this directory to denote the current version of binaries that are installed.
CACHE_VERSIONS := $(CACHE)/versions

# Update the $PATH so we can use buf directly
export PATH := $(abspath $(CACHE_BIN)):$(PATH)
# Update GOBIN to point to CACHE_BIN for source installations
export GOBIN := $(abspath $(CACHE_BIN))
# This is needed to allow versions to be added to Golang modules with go get
export GO111MODULE := on

# BUF points to the marker file for the installed version.
#
# If BUF_VERSION is changed, the binary will be re-downloaded.
BUF := $(CACHE_VERSIONS)/buf/$(BUF_VERSION)
$(BUF):
	$(eval BUF_TMP := $(shell mktemp -d))
	cd $(BUF_TMP); go install github.com/bufbuild/buf/cmd/buf@v$(BUF_VERSION)
	@rm -rf $(BUF_TMP)
	@rm -rf $(dir $(BUF))
	@mkdir -p $(dir $(BUF))
	@touch $(BUF)

# PROTOC points to the marker file for the installed version.
#
# If PROTOC_VERSION is changed, the binary will be re-downloaded.
PROTOC := $(CACHE_VERSIONS)/protoc/$(PROTOC_VERSION)
$(PROTOC):
	@rm -f $(CACHE_BIN)/protoc
	@mkdir -p $(CACHE_BIN)
	$(eval PROTOC_TMP := $(shell mktemp -d))
	curl -sSL \
		"https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-$(PLATFORM)-$(PROTOC_ARCH).zip" \
		-o "$(PROTOC_TMP)/protoc.zip"
	unzip -o "$(PROTOC_TMP)/protoc.zip" -d "$(CACHE)" bin/protoc
	unzip -o "$(PROTOC_TMP)/protoc.zip" -d "$(CACHE)" include/*
	@rm -rf $(PROTOC_TMP)
	chmod +x "$(CACHE_BIN)/protoc"
	@rm -rf $(dir $(PROTOC))
	@mkdir -p $(dir $(PROTOC))
	@touch $(PROTOC)

# PROTOC_GEN_DOC points to the marker file for the installed version.
#
# If PROTOC_GEN_DOC_VERSION is changed, the binary will be re-downloaded.
PROTOC_GEN_DOC := $(CACHE_VERSIONS)/protoc-gen-doc/$(PROTOC_GEN_DOC_VERSION)
$(PROTOC_GEN_DOC):
	@rm -f $(CACHE_BIN)/protoc-gen-doc
	@mkdir -p $(CACHE_BIN)
	$(eval PROTOC_GEN_DOC_TMP := $(shell mktemp -d))
	cd $(PROTOC_GEN_DOC_TMP); go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@v$(PROTOC_GEN_DOC_VERSION)
	@rm -rf $(PROTOC_GEN_DOC_TMP)
	@rm -rf $(dir $(PROTOC_GEN_DOC))
	@mkdir -p $(dir $(PROTOC_GEN_DOC))
	@touch $(PROTOC_GEN_DOC)

# PROTOC_GEN_GO points to the marker file for the installed version.
#
# If PROTOC_GEN_GO_VERSION is changed, the binary will be re-downloaded.
PROTOC_GEN_GO := $(CACHE_VERSIONS)/protoc-gen-go/$(PROTOC_GEN_GO_VERSION)
$(PROTOC_GEN_GO):
	@rm -f $(CACHE_BIN)/protoc-gen-go
	@mkdir -p $(CACHE_BIN)
	$(eval PROTOC_GEN_GO_TMP := $(shell mktemp -d))
	cd $(PROTOC_GEN_GO_TMP); go install google.golang.org/protobuf/cmd/protoc-gen-go@v$(PROTOC_GEN_GO_VERSION)
	@rm -rf $(PROTOC_GEN_GO_TMP)
	@rm -rf $(dir $(PROTOC_GEN_GO))
	@mkdir -p $(dir $(PROTOC_GEN_GO))
	@touch $(PROTOC_GEN_GO)

# PROTOC_GEN_GO_GRPC points to the marker file for the installed version.
#
# If PROTOC_GEN_GO_GRPC_VERSION is changed, the binary will be re-downloaded.
PROTOC_GEN_GO_GRPC := $(CACHE_VERSIONS)/protoc-gen-go-grpc/$(PROTOC_GEN_GO_GRPC_VERSION)
$(PROTOC_GEN_GO_GRPC):
	@rm -f $(CACHE_BIN)/protoc-gen-go-grpc
	@mkdir -p $(CACHE_BIN)
	$(eval PROTOC_GEN_GO_GRPC_TMP := $(shell mktemp -d))
	cd $(PROTOC_GEN_GO_GRPC_TMP); go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v$(PROTOC_GEN_GO_GRPC_VERSION)
	@rm -rf $(PROTOC_GEN_GO_GRPC_TMP)
	@rm -rf $(dir $(PROTOC_GEN_GO_GRPC))
	@mkdir -p $(dir $(PROTOC_GEN_GO_GRPC))
	@touch $(PROTOC_GEN_GO_GRPC)

PROTOC_GEN_GO_GRPC_GATEWAY := $(CACHE_VERSIONS)/protoc-gen-go-grpc-gateway/$(PROTOC_GEN_GO_GRPC_GATEWAY_VERSION)
$(PROTOC_GEN_GO_GRPC_GATEWAY):
	@rm -f $(CACHE_BIN)/protoc-gen-go-grpc-gateway
	@mkdir -p $(CACHE_BIN)
	$(eval PROTOC_GEN_GO_GRPC_GATEWAY_TMP := $(shell mktemp -d))
	cd $(PROTOC_GEN_GO_GRPC_GATEWAY_TMP); go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v$(PROTOC_GEN_GO_GRPC_GATEWAY_VERSION)
	@rm -rf $(PROTOC_GEN_GO_GRPC_GATEWAY_TMP)
	@rm -rf $(dir $(PROTOC_GEN_GO_GRPC_GATEWAY))
	@mkdir -p $(dir $(PROTOC_GEN_GO_GRPC_GATEWAY))
	@touch $(PROTOC_GEN_GO_GRPC_GATEWAY)

PROTOC_GEN_OPENAPIV2 := $(CACHE_VERSIONS)/protoc-gen-openapiv2/$(PROTOC_GEN_GO_GRPC_GATEWAY_VERSION)
$(PROTOC_GEN_OPENAPIV2):
	@rm -f $(CACHE_BIN)/protoc-gen-openapiv2
	@mkdir -p $(CACHE_BIN)
	$(eval PROTOC_GEN_OPENAPIV2_TMP := $(shell mktemp -d))
	cd $(PROTOC_GEN_OPENAPIV2_TMP); go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v$(PROTOC_GEN_GO_GRPC_GATEWAY_VERSION)
	@rm -rf $(PROTOC_GEN_OPENAPIV2_TMP)
	@rm -rf $(dir $(PROTOC_GEN_OPENAPIV2))
	@mkdir -p $(dir $(PROTOC_GEN_OPENAPIV2))
	@touch $(PROTOC_GEN_OPENAPIV2)

WIRE := $(CACHE_VERSIONS)/wire/$(WIRE_VERSION)
$(WIRE):
	@rm -f $(CACHE_BIN)/wire
	@mkdir -p $(CACHE_BIN)
	$(eval WIRE_TMP := $(shell mktemp -d))
	cd $(WIRE_TMP); go install github.com/google/wire/cmd/wire@v$(WIRE_VERSION)
	@rm -rf $(WIRE_TMP)
	@rm -rf $(dir $(WIRE))
	@mkdir -p $(dir $(WIRE))
	@touch $(WIRE)


.PHONY: deps
deps: $(BUF) $(PROTOC) $(PROTOC_GEN_DOC) $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_GRPC) $(PROTOC_GEN_GO_GRPC_GATEWAY) $(PROTOC_GEN_OPENAPIV2) $(WIRE)

.PHONY: cleandeps
cleandeps:
	rm -rf $(CACHE_BASE)