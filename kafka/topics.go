package kafka

import (
	"fmt"
	"github.com/izolight/kafkalib/format"
	"regexp"

	"github.com/Shopify/sarama"
)

// NewTopic holds the information for creating a new topic
type NewTopic struct {
	Name string
	sarama.TopicDetail
}

// CreateTopic creates the topic defined in the TopicClient
func (c Conn) CreateTopic(topic NewTopic) error {
	err := c.AdminClient.CreateTopic(topic.Name, &topic.TopicDetail, false)
	if err != nil {
		return fmt.Errorf("Error creating topic: %s", err)
	}
	return nil
}

// DeleteTopic deletes a topic
func (c Conn) DeleteTopic(topic string) error {
	err := c.AdminClient.DeleteTopic(topic)
	if err != nil {
		return err
	}
	return nil
}
