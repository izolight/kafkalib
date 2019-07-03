package kafka

import (
	"github.com/Shopify/sarama"
	"gitea.izolight.xyz/gabor/kafkalib/format"
)

// BrokerClient is uses for interacting with kafka and saving the data in memory
type BrokerClient struct {
	Client  *sarama.ClusterAdmin
	Brokers []*sarama.Broker
}

// GetBrokers implements the AdminClient interface for BrokerClient
func (c Conn) GetBrokers() ([]*sarama.Broker, error) {
	brokers := c.Client.Brokers()
	return brokers, nil
}

// GetAll implements the AdminClient interface for BrokerClient
func (a BrokerClient) GetAll() (format.Topics, error) {
	panic("not implemented")
}

// Create implements the AdminClient interface for BrokerClient
func (a BrokerClient) Create(topic NewTopic) error {
	panic("not implemented")
}

// Delete implements the AdminClient interface for BrokerClient
func (a BrokerClient) Delete() error {
	panic("not implemented")
}
