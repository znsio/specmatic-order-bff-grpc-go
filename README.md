## Seps to run the project

```
make clean
make all
./specmatic-order-bff-grpc-go
```

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
- [ ] Git submodule for proto
- [ ] Externalising GoPackage name to buf.gen.yaml
- [ ] Architecture diagram
- [ ] Dockerize the project
- [ ] Test Containers
- [ ] Fix buf validate proto file dependency in specmatic-grpc
- [ ] Kafka
