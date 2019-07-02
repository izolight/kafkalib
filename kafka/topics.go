package kafka

import (
	"fmt"
	"gitea.izolight.xyz/gabor/kafkalib/format"
	"regexp"

	"github.com/Shopify/sarama"
)

// NewTopic holds the information for creating a new topic
type NewTopic struct {
	Name string
	sarama.TopicDetail
}

// GetTopic returns the topic defined in the Name of the client
func (c Conn) GetTopic(filter string) (format.Topics, error) {
	allTopics, err := c.AdminClient.ListTopics()
	if err != nil {
		return nil, fmt.Errorf("Error getting allTopics: %s", err)
	}
	topics := format.Topics{}
	r, err := regexp.Compile(filter)
	for k, v := range allTopics {
		if r.MatchString(k) {
			topics[k] = v
		}
	}
	if len(topics) > 0 {
		return topics, nil
	}
	return nil, fmt.Errorf("Topic %s not found", filter)
}

// GetAllTopics returns all known topics
func (c Conn) GetAllTopics() (format.Topics, error) {
	topics, err := c.AdminClient.ListTopics()
	if err != nil {
		return nil, fmt.Errorf("Error getting topics: %s", err)
	}
	return topics, nil
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
