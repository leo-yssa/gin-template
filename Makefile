PROTO_DIR := proto
GO_OUT := grpc/pkg/api
GO_PROTO_PACKAGE := proto
PROTO_FILES := $(wildcard $(PROTO_DIR)/*/*.proto)

WIRE_DIR := ./rest/app/injector
WIRE_DIRS := $(wildcard $(WIRE_DIR)/*/)

generate:
	wire $(WIRE_DIRS)
	protoc -I=$(PROTO_DIR) --go_out=$(GO_OUT) --go-grpc_out=$(GO_OUT) --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative $(PROTO_FILES)
	
build:
	go build -o bin/rest rest/main.go
	go build -o bin/grpc grpc/main.go