package main

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gocql/gocql"
	pb "github.com/ianthpun/micro-example/proto"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/micro/go-micro/broker"
	"log"
)

type Processor struct {
	DAL *gocql.Session
}

func NewProcessor(dal *gocql.Session) *Processor {
	return &Processor{
		DAL: dal,
	}
}

func (p *Processor) process(messages <-chan *message.Message) {
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
