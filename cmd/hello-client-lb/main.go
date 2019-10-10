package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"

	"github.com/ellulpatrick/grpc-demo/pkg/grpc-hello/v2"
)

func main() {

	r, cleanup := registerResolver()
	defer cleanup()

	conn, err := grpc.Dial(
		"dummy",
		grpc.WithInsecure(), // because we're not using TLS
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
		grpc.WithBalancerName(roundrobin.Name),
	)
	if err != nil {
		log.Fatalf("Could not set up connection: %v", err)
	}

	addrs := os.Getenv("SERVERS")
	setAddrs(r, addrs)

	log.Printf("We have a magical connection object to %v...", addrs)

	helloClient := grpc_hello.NewHelloServiceClient(conn)

	go loop("Bluey", 4*time.Second, 6, helloClient)
	time.Sleep(3 * time.Second) // to offset
	loop("Bingo", 10*time.Second, 0, helloClient)
}

func loop(name string, clientSleep time.Duration, serverDelay int64, helloClient grpc_hello.HelloServiceClient) {
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
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
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
