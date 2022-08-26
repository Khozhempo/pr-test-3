package kafka

import (
	"github.com/Shopify/sarama"
)

type Opts struct {
	Address []string
	//Port    string
	Topic string
}

type RepositoryKafkaStore struct {
	kafkaConn  *sarama.AsyncProducer
	kafkaTopic *string
}

func NewStore(kafkaConn *sarama.AsyncProducer, kafkaTopic *string) *RepositoryKafkaStore {
	return &RepositoryKafkaStore{
		kafkaConn:  kafkaConn,
		kafkaTopic: kafkaTopic,
	}
}
