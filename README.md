## Steps to run the project

```
go mod download
make clean
make all
./specmatic-order-bff-grpc-go
```

## Running Contract Tests

* Start stub server - `java -jar specmatic-grpc.jar stub`
* Run the app - `./specmatic-order-bff-grpc-go`
* Run contract tests (with API resiliency switched on) - `java -DSPECMATIC_GENERATIVE_TESTS=true -jar specmatic-grpc.jar test`

## Debugging steps

In case of pb files already exists,
```
rm -rf ~/Library/Caches/buf
```

## Dev notes - Ignore

```
brew install go
brew install protobuf
brew install protoc
brew install bufbuild/buf/buf
brew install protoc-gen-go-grpc
go mod tidy
buf dep update
```

## TODO

- [ ] Generative Tests
  - [x] Update specmatic-grpc to run all tests even when there are failures
  - [x] Add validations
  - [x] Leverage protovalidate?
- [ ] Git submodule for proto
- [ ] Externalising GoPackage name to buf.gen.yaml
- [ ] Architecture diagram
- [ ] Run everything in Github CI
- [ ] Dockerize the project
- [ ] Test Containers
- [ ] Fix buf validate proto file dependency in specmatic-grpc
- [ ] Kafka
