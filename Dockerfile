# Build stage
FROM golang:1.22-alpine AS builder
# FROM bufbuild/buf:latest AS builder

# Install necessary tools
RUN apk add --no-cache make && \
    apk add --no-cache curl && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    # go install github.com/bufbuild/buf/cmd/buf@latest
    curl -L -o /usr/local/bin/buf https://github.com/bufbuild/buf/releases/latest/download/buf-Linux-x86_64 && \
    chmod +x /usr/local/bin/buf


WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the rest of the source code and Makefile
COPY . .

# Build the application
RUN make all

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/specmatic-order-bff-grpc-go .

# Expose the port your service runs on
EXPOSE 8080

# Run the binary
CMD ["./specmatic-order-bff-grpc-go"]