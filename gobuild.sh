#!/bin/bash

go build -o bin/hello-client1 cmd/hello-client1/*.go
go build -o bin/hello-client2 cmd/hello-client2/*.go
go build -o bin/hello-server1 cmd/hello-server1/*.go
go build -o bin/hello-server2 cmd/hello-server2/*.go
go build -o bin/hello-client-lb cmd/hello-client-lb/*.go
