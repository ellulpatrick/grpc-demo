# A little gRPC demo in Go

## Running the Servers
```
LISTEN=:8091 go run cmd/hello-server/main.go
```
and
```
LISTEN=:8092 go run cmd/hello-server2/main.go
```

## Running the Clients

## How to compile proto files

The compiled go files are checked in, but if you need to recompile them.

You need to have the proto compiler (`protoc`) installed with the go plugins

Then, from the `pkg` directory:

```
protoc --go_out=plugins=grpc,paths=source_relative:. grpc-hello/v1/grpc-hello.proto
```
and for the v2
```
protoc --go_out=plugins=grpc,paths=source_relative:. grpc-hello/v2/grpc-hello.proto
```

# Other Links
