package kafka

import (
	"errors"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

func (opt *Opts) createTopic(config *sarama.Config) (err error) {
	conn := sarama.NewBroker(opt.Address[0])
	err = conn.Open(config)
	if err != nil {
		return err
	}

	defer conn.Close()

	_, err = conn.Connected()
	if err != nil {
		return err
	}

	topicDetail := &sarama.TopicDetail{}
	topicDetail.NumPartitions = int32(1)
	topicDetail.ReplicationFactor = int16(1)
	topicDetail.ConfigEntries = make(map[string]*string)
	topicDetails := make(map[string]*sarama.TopicDetail)
	topicDetails[opt.Topic] = topicDetail

	request := sarama.CreateTopicsRequest{
		Timeout: time.Second * 15,
	}

	_, err = conn.CreateTopics(&request)
	if err != nil {
		return err
	}

	return nil
}

func checkInStringSlice(searchIn []string, lookFor string) bool {
	for _, value := range searchIn {
		if value == lookFor {
			return true
		}
	}
	return false
}

func (kafka *RepositoryKafkaStore) StartListener() {
	go func(p sarama.AsyncProducer) {
		for {
			select {
			case <-p.Successes():
				// On success
			case fail := <-p.Errors():
				log.Println("err: ", fail.Err)
			}
		}
	}(*kafka.kafkaConn)
}

func InitKafka() (*sarama.AsyncProducer, *string, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 2 * time.Second

	kafkaOpts, err := GetKafkaConfig()
	if err != nil {
		return nil, nil, err
	}

	err = kafkaOpts.CheckIfTopicExist(config)
	if err != nil {
		return nil, nil, err
	}

	kafkaConn, err := sarama.NewAsyncProducer(kafkaOpts.Address, config)
	if err != nil {
		return nil, nil, errors.New("Error occurred while establishing connection to Kafka " + err.Error())
	}

	return &kafkaConn, &kafkaOpts.Topic, nil
}

func addPortToAddr(addr []string, port string) []string {
	for k, v := range addr {
		addr[k] = v + ":" + port
	}
	return addr
}
