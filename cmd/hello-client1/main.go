package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/grpc"

	"github.com/ellulpatrick/grpc-demo/pkg/grpc-hello/v1"
	"time"
)

func main() {

	serverAddr := os.Getenv("SERVER")
	conn, err := grpc.Dial(
		serverAddr,
		grpc.WithInsecure(), // because we're not using TLS
	)

	if err != nil {
		log.Fatalf("Could not set up connection: %v", err)
	}

	log.Printf("We have a magical connection object to %v...", serverAddr)

	helloClient := grpc_hello.NewHelloServiceClient(conn)

	go loop("Green", 5 * time.Second, helloClient)
	time.Sleep(3 * time.Second) // to offset
	loop("Yellow", 10 * time.Second, helloClient)
}

func loop(name string, clientSleep time.Duration, helloClient grpc_hello.HelloServiceClient) {
	for {
		err := callServer(name, helloClient)

		if err != nil {
			errorSleep := 5 * time.Second
			log.Printf("%v: Sleeping %v After Error", name, errorSleep)
			time.Sleep(errorSleep)
			continue
		}

		log.Printf("%v: Sleeping %v", name, clientSleep)
		time.Sleep(clientSleep)
	}
}

func callServer(name string, helloClient grpc_hello.HelloServiceClient) error {

	log.Printf("%v: Calling Server.", name)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 15 * time.Second)
	defer cancel()

	response, err := helloClient.SayHello(ctx,
		&grpc_hello.SayHelloRequest{
			Neme:  name,
		})

	if err != nil {
		log.Printf("%v: Received Error: %v", name, err)
		return err
	}

	greeting := response.Greeting
	log.Printf("%v: Received greeting: %v", name, greeting)
	return nil
}