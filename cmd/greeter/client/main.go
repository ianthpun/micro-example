package main

import (
	"context"
	"fmt"
	"time"

	proto "github.com/ianthpun/micro-example/proto"
	micro "github.com/micro/go-micro"
)

func main() {
	// Create a new service
	service := micro.NewService(micro.Name("greeter.client"))
	// Initialise the client and parse command line flags
	service.Init()

	// Create new greeter client
	greeter := proto.NewGreeterService("greeter", service.Client())

	// Call the greeter
	for {
		time.Sleep(5 * time.Second)
		rsp, err := greeter.Hello(context.Background(), &proto.Request{Name: "John"})
		if err != nil {
			panic(err)
		}

		// Print response
		fmt.Println(rsp.Msg)
	}
}
