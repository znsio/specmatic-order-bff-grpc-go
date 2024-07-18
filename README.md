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

### Contract Testing BFF using specmatic-grpc docker image (preferred) with test containers

```shell
go mod tidy
go test contract_test.go -v -count=1 
```

### Contract Testing BFF by building and running the application using specmatic-grpc JAR file (old-school)

1. Installing pre-requisites

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
   1.1. Add protoc-gen-go to your path using the following command:

```shell
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bash_profile 
source ~/.bash_profile 
```   

2. Installing the `buf` tool
```shell
brew install bufbuild/buf/buf
```
For more options to install buf tool, refer to [buf installation guide](https://docs.buf.build/installation)

3. Building the application

```shell
go mod tidy
go mod download
make clean
make all
```

## Running Contract Tests

### Using specmatic-grpc JAR file

* Start Specmatic stub server - 
```shell
java -jar lib/specmatic-grpc-0.0.3-TRIAL-all.jar stub
```
* Run the Go BFF app
```shell
./specmatic-order-bff-grpc-go
```
* Run contract tests (with API resiliency switched on)
```shell
java -DSPECMATIC_GENERATIVE_TESTS=true -jar lib/specmatic-grpc-0.0.3-TRIAL-all.jar test --port=8080
```

## Debugging steps

In case of pb files already exists,
```shell
rm -rf ~/Library/Caches/buf
```
