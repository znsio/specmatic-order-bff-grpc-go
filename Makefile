PROTO_DIR := ./proto_files
GO_OUT_DIR := ./pkg/api
CMD_DIR := ./cmd
BFF_SERVICE_BINARY := bff_service

.PHONY: all proto build clean

all: proto build

# Generate Go code from proto files using buf
proto:
	@echo "Generating Go code from proto files using buf..."
	@buf generate

# Build the application
build: proto
	@echo "Building the application..."
	@go build -o $(BFF_SERVICE_BINARY) $(CMD_DIR)/main.go

# Clean up generated files and binaries
clean:
	@echo "Cleaning up..."
	@rm -rf $(GO_OUT_DIR)/*
	@rm -f $(BFF_SERVICE_BINARY)