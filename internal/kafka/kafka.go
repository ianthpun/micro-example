package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	watermill_kafka "github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
)

func GetSubscriber(confOpts *sarama.Config, brokers []string) (*watermill_kafka.Subscriber, error) {
	conf := confOpts
	if confOpts == nil {
		conf = watermill_kafka.DefaultSaramaSubscriberConfig()
		// equivalent of auto.offset.reset: earliest
		conf.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	return watermill_kafka.NewSubscriber(
		watermill_kafka.SubscriberConfig{
			Brokers:               brokers,
			Unmarshaler:           watermill_kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: conf,
			ConsumerGroup:         "test_consumer_group",
		},
		watermill.NewStdLogger(false, false),
	)
}
