## Architecture

![Specmatic gRPC Support Architecture](/assets/SpecmaticGRPCSupport.gif)

## Setup

### Cloning with submodules

1. Clone the repository

   ```shell
   git clone https://github.com/znsio/specmatic-order-bff-grpc-go
   ```

2. Initialize and update the `proto_files` submodule

   ```shell
   git submodule update --init --recursive --remote
   ```

3. Enable automatic submodule updating when executing `git pull`

   ```shell
   git config submodule.recurse true
   ```

### Building and running the application

1. Installing prequisites

    ```shell
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    ```

2. Building the application

   ```
   go mod tidy
   go mod download
   make clean
   make all
   ./specmatic-order-bff-grpc-go
   ```
   

## Running Contract Tests

* Start stub server - `java -jar specmatic-grpc.jar stub`
* Run the app - `./specmatic-order-bff-grpc-go`
* Run contract tests (with API resiliency switched on) - `java -DSPECMATIC_GENERATIVE_TESTS=true -jar specmatic-grpc.jar test --port=50051`

## Debugging steps

In case of pb files already exists,
```
rm -rf ~/Library/Caches/buf
```
