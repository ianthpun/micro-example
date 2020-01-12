package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-log/log"
	"github.com/google/uuid"
	"github.com/ianthpun/micro-example/internal/kafka"
	proto "github.com/ianthpun/micro-example/proto"
	micro "github.com/micro/go-micro"
)

func processEvent (ctx context.Context, event proto.GreetedEvent) error {
	fmt.Printf("Receied the roundtrip: %+v", event)
	return nil
}

func main() {

	// Create a new service
	service := micro.NewService(micro.Name("greeter.client"))
	// Initialise the client and parse command line flags
	service.Init()
	// create publisher
	pub1 := micro.NewPublisher(kafka.GREETING_TOPIC, service.Client())

	// Create new greeter client
	//greeter := proto.NewGreeterService("greeter", service.Client())

	micro.RegisterSubscriber(kafka.GREETING_TOPIC, service.Server(), processEvent)

	for {
		time.Sleep(5 * time.Second)
		ev := &proto.GreetEvent{
			Id:  uuid.New().String(),
			Msg: fmt.Sprintf("Im Jon through kafka"),
		}

		fmt.Printf("publishing %+v\n", ev)

		// publish an event
		if err := pub1.Publish(context.Background(), ev); err != nil {
			log.Logf("error publishing: %v", err)
		}

	}
}
