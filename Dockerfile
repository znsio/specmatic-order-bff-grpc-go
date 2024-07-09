FROM golang:latest AS builder

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install github.com/bufbuild/buf/cmd/buf@latest

# Set up the working directory
WORKDIR /app

# Copy the source code and Makefile
COPY . .

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