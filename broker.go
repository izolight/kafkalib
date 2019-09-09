package kafkalib

import "github.com/Shopify/sarama"

type Broker struct {
	ID   string `json:"id"`
	Addr string `json:"addr"`
	Rack string `json:"rack"`
}

type Brokers []Broker

// GetBrokers implements the AdminClient interface for BrokerClient
func (c Conn) GetBrokers() ([]Broker, error) {
	brokers := c.Client.Brokers()
	b := Brokers{}
	err := b.UnmarshallSarama(brokers)
	return b, err
}

func (b Broker) MarshallSarama() (*sarama.Broker, error) {

	return nil, nil
}

func (b Broker) UnmarshallSarama(broker *sarama.Broker) error {
	b.ID = string(broker.ID())
	b.Addr = broker.Addr()
	b.Rack = broker.Rack()

	return nil
}

func (b Brokers) UnmarshallSarama(brokers []*sarama.Broker) error {
	for _, broker := range brokers {
		br := Broker{}
		err := br.UnmarshallSarama(broker)
		if err != nil {
			return err
		}
		b = append(b, br)
	}
	return nil
}