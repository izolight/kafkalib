package kafkalib

import "github.com/Shopify/sarama"

type Marshaler interface {
	MarshalSarama() ([]*sarama.ResourceAcls, error)
}

type Unmarshaler interface {
	UnmarshalSarama([]*sarama.ResourceAcls) error
}