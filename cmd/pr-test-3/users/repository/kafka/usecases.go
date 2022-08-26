package kafka

import (
	"encoding/json"
	"errors"
	"github.com/Shopify/sarama"
	"os"
	"strings"
	"time"
)

type Message struct {
	Cmd       string `json:"cmd"`
	Username  string `json:"username"`
	Timestamp int64  `json:"timestamp"`
}

func GetKafkaConfig() (Opts, error) {
	var outArgs Opts

	port := os.Getenv("TEST_KAFKA_PORT")
	if port == "" {
		return Opts{}, errors.New("TEST_KAFKA_PORT is not set")
	}
	//outArgs.Port = port

	topic := os.Getenv("TEST_KAFKA_TOPIC")
	if topic == "" {
		return Opts{}, errors.New("TEST_KAFKA_TOPIC is not set")
	}
	outArgs.Topic = topic

	address := os.Getenv("TEST_KAFKA_ADDR")
	if address == "" {
		return Opts{}, errors.New("TEST_KAFKA_ADDR is not set")
	}
	outArgs.Address = strings.Split(address, ",")
	outArgs.Address = addPortToAddr(outArgs.Address, port)

	return outArgs, nil
}

func (opt *Opts) CheckIfTopicExist(config *sarama.Config) error {
	conn, err := sarama.NewConsumer(opt.Address, config)
	if err != nil {
		return err
	}

	defer conn.Close()

	topics, _ := conn.Topics()
	if !checkInStringSlice(topics, opt.Topic) {
		if err = opt.createTopic(config); err != nil {
			return err
		}
	}

	return nil
}

func (kafka *RepositoryKafkaStore) LogNewUser(username string) error {
	var msg sarama.ProducerMessage

	msg.Topic = *kafka.kafkaTopic
	msg.Key = sarama.StringEncoder("pr-test-3")
	val, _ := json.Marshal(&Message{Cmd: "AddNewUser", Username: username, Timestamp: time.Now().Unix()})

	msg.Value = sarama.ByteEncoder(val)
	msg.Timestamp = time.Now()

	conn := *kafka.kafkaConn
	conn.Input() <- &msg

	return nil
}
