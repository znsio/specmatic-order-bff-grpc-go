# # Use Ubuntu as the base image
# FROM ubuntu:22.04

# # Avoid prompts from apt
# ENV DEBIAN_FRONTEND=noninteractive

# # Install necessary tools
# RUN apt-get update && apt-get install -y \
#     wget \
#     git \
#     make \
#     gcc \
#     && rm -rf /var/lib/apt/lists/*

# # Install Go 1.22.5
# RUN wget https://go.dev/dl/go1.22.5.linux-amd64.tar.gz \
#     && tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz \
#     && rm go1.22.5.linux-amd64.tar.gz

# # Set Go environment variables
# ENV PATH="/usr/local/go/bin:${PATH}"
# ENV GOPATH="/go"
# ENV PATH="${GOPATH}/bin:${PATH}"

# # Install Go tools
# RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# RUN go install github.com/bufbuild/buf/cmd/buf@latest

# # Set up the working directory
# WORKDIR /app

# # Copy the source code and Makefile
# COPY . .

# # Download all dependencies
# RUN go mod tidy
# RUN go mod download

# # Run make commands
# RUN make clean
# RUN make all SHELL='sh -x'

# # Expose the port your service runs on
# EXPOSE 8080

# # Run the binary
# CMD ["./specmatic-order-bff-grpc-go"]

FROM golang:latest AS builder

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install github.com/bufbuild/buf/cmd/buf@latest

# Set up the working directory
WORKDIR /app

# Copy the source code and Makefile
COPY ../../../../../. .

# # Copy go mod and sum files
# COPY go.mod go.sum ./

# Download all dependencies
RUN go mod tidy
RUN go mod download

# Copy the source code and Makefile
COPY . .

RUN make all

# # Final stage
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates

# WORKDIR /root/

# # Copy the binary from the builder stage
# COPY --from=builder /app/specmatic-order-bff-grpc-go .

# Expose the port your service runs on
EXPOSE 8080

# Run the binary
CMD ["./specmatic-order-bff-grpc-go"]