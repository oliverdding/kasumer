package main

import (
	"fmt"
	"testing"

	"github.com/segmentio/kafka-go"
)

func Test_showTopics(t *testing.T) {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	for k := range m {
		fmt.Println(k)
	}
}

func Test_createTopic(t *testing.T) {
	// to create topics when auto.create.topics.enable='false'
	topic := "test"
	partition := 3

	controllerConn, err := kafka.Dial("tcp", "127.0.0.1:9092")
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     partition,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}

}
