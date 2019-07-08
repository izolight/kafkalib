package kafkalib

import "github.com/Shopify/sarama"

type Marshaler interface {
	MarshalResourceAcls() ([]*sarama.ResourceAcls, error)
}

type Unmarshaler interface {
	UnmarshalResourceAcls([]*sarama.ResourceAcls) error
}