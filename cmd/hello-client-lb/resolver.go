package main

import (
	"strings"

	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
)

func registerResolver() (*manual.Resolver, func()) {

	r, cleanup := manual.GenerateAndRegisterManualResolver()

	// Each resolver scheme package calls init() to register
	// itself. So we need to call Register() again here to
	// register the selected builder.
	resolver.Register(r)
	resolver.SetDefaultScheme(r.Scheme())

	return r, cleanup
}

func setAddrs(r *manual.Resolver, serverIPs string) {
	addresses := []resolver.Address{}
	for _, addr := range strings.Split(serverIPs, ",") {
		addresses = append(addresses, resolver.Address{Addr: addr, Type: resolver.Backend})
	}

	r.UpdateState(resolver.State{Addresses:addresses})
}