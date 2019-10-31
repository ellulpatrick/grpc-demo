# A little gRPC demo in Go

## Running the Servers
```
LISTEN=:8091 go run cmd/hello-server1/main.go
```
and
```
LISTEN=:8092 go run cmd/hello-server2/main.go
```

## Running the Clients
```
SERVER=:8091 go run cmd/hello-client1/main.go
```
and
```
SERVER=:8092 go run cmd/hello-client2/main.go
```
and
```
SERVERS=:8091,:8092 go run cmd/hello-client-lb/main.go
```

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
- https://grpc.io/ 
- https://developers.google.com/protocol-buffers/docs/proto3
- https://developers.google.com/protocol-buffers/docs/proto3#updating
- https://kubernetes.io/blog/2018/11/07/grpc-load-balancing-on-kubernetes-without-tears/ 
- https://itnext.io/on-grpc-load-balancing-683257c5b7b3 