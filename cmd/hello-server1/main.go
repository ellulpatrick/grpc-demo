package main

import (
	"net"
	"log"
	"os"
	"syscall"
	"os/signal"
	"fmt"
	"context"

	"google.golang.org/grpc"

	"github.com/ellulpatrick/grpc-demo/pkg/grpc-hello/v1"
)

type helloService struct{}

func (hs *helloService) SayHello(ctx context.Context, request *grpc_hello.SayHelloRequest) (*grpc_hello.SayHelloResponse, error) {
	log.Printf("Received from %v, returning.", request.Neme)
	return &grpc_hello.SayHelloResponse{
		Greeting: fmt.Sprintf("Hello %v", request.Neme),
	}, nil
}

func main() {

	grpcSvr := grpc.NewServer()

	helloSvc := &helloService{}

	grpc_hello.RegisterHelloServiceServer(grpcSvr, helloSvc)

	go func() {
		listenAddr := os.Getenv("LISTEN")
		listener, err := net.Listen("tcp", listenAddr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		log.Printf("grpcSrv V1 Listening on %v", listenAddr)
		err = grpcSvr.Serve(listener) // Blocks
		log.Printf("grpcSvr returned. Err: %v", err)
	}()

	exitCh := make(chan os.Signal, 1)
	signal.Notify(exitCh, os.Interrupt, syscall.SIGTERM)
	<-exitCh

	log.Printf("Stopping...")
	grpcSvr.GracefulStop()
}

