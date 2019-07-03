package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"gitea.izolight.xyz/gabor/kafkalib/format"
)

// ConsumerGroupClient is uses for interacting with kafka and saving the data in memory
type ConsumerGroupClient struct {
	Client         sarama.ClusterAdmin
	ConsumerGroups []*sarama.ConsumerGroup
}

// GetConsumerGroup returns a single Consumer Group according to filter
func (c Conn) GetConsumerGroup() (format.Topics, error) {
	panic("not implemented")
}

// GetAllConsumerGroups returns all Consumer Groups
func (c Conn) GetAllConsumerGroups() (format.ConsumerGroups, error) {
	groups, err := c.AdminClient.ListConsumerGroups()
	if err != nil {
		return nil, fmt.Errorf("Error getting all Consumergroups: %s", err)
	}
	return groups, nil
}

// CreateConsumerGroup creates a new Consumer Group
func (c Conn) CreateConsumerGroup(topic NewTopic) error {
	panic("not implemented")
}

// DeleteConsumerGroup deletes a Consumer Group according to filter
func (c Conn) DeleteConsumerGroup() error {
	panic("not implemented")
}
