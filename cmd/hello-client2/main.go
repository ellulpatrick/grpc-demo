package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/grpc"

	"github.com/ellulpatrick/grpc-demo/pkg/grpc-hello/v2"
	"time"
)

func main() {

	serverAddr := os.Getenv("SERVER")
	conn, err := grpc.Dial(
		serverAddr,
		grpc.WithInsecure(), // because we're not using TLS
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	)

	if err != nil {
		log.Fatalf("Could not set up connection: %v", err)
	}

	log.Printf("We have a magical connection objectto %v...", serverAddr)

	helloClient := grpc_hello.NewHelloServiceClient(conn)

	go loop("Bluey", 4 * time.Second, 6, helloClient)
	time.Sleep(1 * time.Second) // to offset
	loop("Bingo", 2 * time.Second, 8, helloClient)
}


func loop(name string, clientSleep time.Duration,  serverDelay int64, helloClient grpc_hello.HelloServiceClient) {
	for {
		err := callServer(name, serverDelay, helloClient)

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

func callServer(name string, delay int64, helloClient grpc_hello.HelloServiceClient) error {

	log.Printf("%v: Calling Server. Delay %v", name, delay)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 30 * time.Second)
	defer cancel()

	response, err := helloClient.SayHello(ctx,
		&grpc_hello.SayHelloRequest{
			Name:  name,
			Delay: delay,
		})

	if err != nil {
		log.Printf("%v: Received Error: %v", name, err)
		return err
	}

	greeting := response.Greeting
	if greeting == "" {
		greeting = "OBS: " + response.OBSOLETEGreeting
	}
	log.Printf("%v: Received greeting: %v", name, greeting)
	return nil
}