package main

import (
	"context"
	"github.com/ianthpun/micro-example/internal/cassandra"
	"github.com/ianthpun/micro-example/internal/kafka"
)


func main() {
	cql, err := cassandra.Connect(cassandra.Config{
		Host:     "127.0.0.1",
		Keyspace: "testplayground",
	})
	if err != nil {
		panic(err)
	}
	defer cql.Close()

	subscriber, err := kafka.GetSubscriber(nil,[]string{"kafka:9092"})
	if err != nil {
		panic(err)
	}
	defer subscriber.Close()

	messages, err := subscriber.Subscribe(context.Background(), kafka.GREETING_TOPIC)
	if err != nil {
		panic(err)
	}

	processor := NewProcessor(cql)
	processor.process(messages)
}

