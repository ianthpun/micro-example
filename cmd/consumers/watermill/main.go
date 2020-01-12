package main

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/broker"
	"log"

	"github.com/Shopify/sarama"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gogo/protobuf/proto"
	pkgkafka "github.com/ianthpun/micro-example/internal/kafka"
	pb "github.com/ianthpun/micro-example/proto"
)

func main() {
	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
	// equivalent of auto.offset.reset: earliest
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               []string{"kafka:9092"},
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: saramaSubscriberConfig,
			ConsumerGroup:         "test_consumer_group",
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	messages, err := subscriber.Subscribe(context.Background(), pkgkafka.GREETING_TOPIC)
	if err != nil {
		panic(err)
	}

	process(messages)

}

func process(messages <-chan *message.Message) {
	for msg := range messages {
		//Incoming messages from this topic are in `broker.Message` style, given from go-micro
		var m broker.Message
		if err := json.Unmarshal(msg.Payload, &m); err != nil {
			panic(err)
		}
		var greetingEvent pb.GreetEvent
		if err := proto.Unmarshal(m.Body, &greetingEvent); err != nil {
			log.Print(err)
		}

		log.Printf("received message: %+v",greetingEvent.Msg)
		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}
