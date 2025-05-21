PROTO_DIR := ./
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)
GO_OUT := .

all: generate

generate:
	@echo "Generating Go files from proto (without gRPC)..."
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(GO_OUT) --go_opt=paths=source_relative \
		$(PROTO_FILES)

.PHONY: all generate