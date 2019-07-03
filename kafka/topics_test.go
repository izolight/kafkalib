package kafka_test

import (
	"github.com/izolight/kafkalib/kafka"
	"testing"

	"github.com/Shopify/sarama"
)

func TestTopic_GetAll(t *testing.T) {
	client := NewTestClient()
	testCases := []struct {
		success bool
	}{
		{true},
		{false},
	}
	for _, tc := range testCases {
		c := kafka.Conn{
			AdminClient: client,
		}
		if !tc.success {
			client.Close()
		}
		topics, err := c.GetAllTopics()
		if err != nil && tc.success {
			t.Fatal(err)
		}
		if topics == nil && tc.success {
			t.Fatal("Didn't return any topics")
		}
		if topics != nil && !tc.success {
			t.Fatal("Returned topic even though it shouldn't")
		}
	}
}

func TestTopic_Get(t *testing.T) {
	client := NewTestClient()
	testCases := []struct {
		name    string
		exists  bool
		success bool
	}{
		{"simpleTopic", true, true},
		{"can't find topic", false, true},
		{"simpleTopic", true, false},
	}
	for _, tc := range testCases {
		c := kafka.Conn{
			AdminClient: client,
		}
		if !tc.success {
			client.Close()
		}
		topic, err := c.GetTopic(tc.name)
		if err != nil && tc.exists && tc.success {
			t.Fatal(err)
		}
		if topic != nil && !tc.success {
			t.Fatal("Returned topic even though it shouldn't")
		}
	}
}

func TestTopic_Create(t *testing.T) {
	client := NewTestClient()
	testCases := []struct {
		name       string
		partitions int32
		replicas   int16
		create     bool
		success    bool
	}{
		{"newTopic", 2, 1, true, true},
		{"anotherNewTopic", 1, 3, false, true},
		{"newTopic", 2, 1, true, false},
	}
	for _, tc := range testCases {
		c := kafka.Conn{
			AdminClient: client,
		}
		if tc.create {
			newTopic := kafka.NewTopic{
				tc.name,
				sarama.TopicDetail{
					NumPartitions:     tc.partitions,
					ReplicationFactor: tc.replicas,
				},
			}
			err := c.CreateTopic(newTopic)
			if err != nil && tc.success {
				t.Fatal(err)
			}
		}
		topics, err := c.GetAllTopics()
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := topics[tc.name]; !ok && tc.create {
			t.Fatalf("Topic %s was not found after creation", tc.name)
		}
	}
}

func TestTopic_Delete(t *testing.T) {
	client := NewTestClient()
	testCases := []struct {
		name   string
		create bool
	}{
		{"newTopic", true},
		{"newTopicDontCreate", false},
	}
	for _, tc := range testCases {
		c := kafka.Conn{
			AdminClient: client,
		}
		if tc.create {
			newTopic := kafka.NewTopic{
				tc.name,
				sarama.TopicDetail{
					NumPartitions:     1,
					ReplicationFactor: 1,
				},
			}
			err := c.CreateTopic(newTopic)
			if err != nil {
				t.Fatal(err)
			}
		}
		err := c.DeleteTopic(tc.name)
		if err != nil && tc.create {
			t.Fatal(err)
		}
		_, err = c.GetTopic(tc.name)
		if err == nil {
			t.Fatalf("Found topic %s, even though if should be deleted", tc.name)
		}
	}
}
