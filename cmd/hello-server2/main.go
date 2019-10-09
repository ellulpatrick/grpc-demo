package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"

	"github.com/ellulpatrick/grpc-demo/pkg/grpc-hello/v2"
)

type helloService struct{}

func (hs *helloService) SayHello(ctx context.Context, request *grpc_hello.SayHelloRequest) (*grpc_hello.SayHelloResponse, error) {

	delay := time.Second * time.Duration(request.Delay)
	log.Printf("Received from %v, first sleeping %v", request.Name, delay)
	time.Sleep(delay)
	log.Printf("Returning to %v, after sleeping %v", request.Name, delay)

	return &grpc_hello.SayHelloResponse{
		OBSOLETEGreeting: fmt.Sprintf("Hello %v", request.Name),
		Greeting:         fmt.Sprintf("Hello %v after %v", request.Name, delay),
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
		log.Printf("grpcSrv V2 Listening on %v", listenAddr)
		err = grpcSvr.Serve(listener) // Blocks
		log.Printf("grpcSvr returned. Err: %v", err)
	}()

	exitCh := make(chan os.Signal, 1)
	signal.Notify(exitCh, os.Interrupt, syscall.SIGTERM)
	<-exitCh

	log.Printf("Stopping...")
	grpcSvr.GracefulStop()
}
