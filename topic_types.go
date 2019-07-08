package kafkalib

import (
	"github.com/Shopify/sarama"
	"github.com/izolight/kafkalib/format"
)

// Topic represents the serialized Topic
type Topic struct {
	Name              string `json:"name"`
	Partitions        int    `json:"partitions"`
	ReplicationFactor int    `json:"replication_factor"`
	RetentionMs       int    `json:"retention_ms,omitempty"`
	ACLs              []ACLs `json:"acls,omitempty"`
}

// TODO: consolidate with format package

func isConsumer(operations []format.Operation) bool {
	read := false
	describe := false
	for _, o := range operations {
		if sarama.AclOperation(o) == sarama.AclOperationRead {
			read = true
		} else if sarama.AclOperation(o) == sarama.AclOperationDescribe {
			describe = true
		}
	}
	return read && describe
}

func isProducer(operations []format.Operation) bool {
	write := false
	describe := false
	create := false
	for _, o := range operations {
		if sarama.AclOperation(o) == sarama.AclOperationWrite {
			write = true
		} else if sarama.AclOperation(o) == sarama.AclOperationDescribe {
			describe = true
		} else if sarama.AclOperation(o) == sarama.AclOperationCreate {
			create = true
		}
	}
	return write && describe && create
}