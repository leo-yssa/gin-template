PROTO_DIR := proto
GO_OUT := grpc/pkg/api
GO_PROTO_PACKAGE := proto
PROTO_FILES := $(wildcard $(PROTO_DIR)/*/*.proto)

WIRE_DIR := ./rest/app/injector
WIRE_DIRS := $(wildcard $(WIRE_DIR)/*/)

include gotools.mk

.PHONY: gotools
gotools: deps

.PHONY: genprotos
genprotos:
	protoc -I=$(PROTO_DIR) --go_out=$(GO_OUT) --go-grpc_out=$(GO_OUT) --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative $(PROTO_FILES)

.PHONY: genwire
genwire:
	wire $(WIRE_DIRS)

.PHONY: build
build:
	go build -o bin/rest rest/main.go
	go build -o bin/grpc grpc/main.go