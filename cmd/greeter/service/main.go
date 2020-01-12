package main

import (
	"context"
	"fmt"

	"github.com/ianthpun/micro-example/internal/kafka"
	proto "github.com/ianthpun/micro-example/proto"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
)

type Greeter struct{}

func (g *Greeter) GreetingCommand(ctx context.Context, req *proto.Command, rsp *proto.Response) error {
	rsp.Msg = req.Msg
	return nil
}
func processEvent(ctx context.Context, event *proto.GreetEvent) error {
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[greeting.topic] Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
	)

	// Init will parse the command line flags.
	service.Init()

	micro.RegisterSubscriber(kafka.GREETING_TOPIC, service.Server(), processEvent)

	// Register handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
