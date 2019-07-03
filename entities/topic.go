package entities

import (
	"github.com/Shopify/sarama"
	"github.com/izolight/kafkalib/format"
	"strconv"
)

// Topic represents the serialized Topic
type Topic struct {
	Name              string   `json:"name"`
	Present           state    `json:"state"`
	Partitions        int      `json:"partitions"`
	ReplicationFactor int      `json:"replication-factor"`
	RetentionMs       int      `json:"retention.ms,omitempty"`
	ACL               TopicACL `json:"acl,omitempty"`
}

// TopicACL represents the serialized ACLs for a Topic
type TopicACL struct {
	Producer []TopicUser `json:"producer,omitempty"`
	Consumer []TopicUser `json:"consumer,omitempty"`
}

// TopicUser represents the serialized User Principal for a Topic
type TopicUser struct {
	Principal string `json:"principal"`
	Present   state  `json:"state"`
}

// ConvertToSerialisableTopic is used to convert a sarama Topic to a serialisable format
func ConvertToSerialisableTopic(t sarama.TopicDetail, acl TopicACL) (Topic, error) {
	topic := Topic{
		Partitions:        int(t.NumPartitions),
		ReplicationFactor: int(t.ReplicationFactor),
		Present:           true,
	}
	retention, ok := t.ConfigEntries["retention.ms"]
	if ok {
		ret, err := strconv.Atoi(*retention)
		if err != nil {
			return topic, err
		}
		topic.RetentionMs = ret
	}

	topic.ACL = acl

	return topic, nil
}

// ConvertToSerialisableACL is used to convert a sarama ACL to a serialisable acl
func ConvertToSerialisableACL(acls []format.ACL) (TopicACL, error) {
	aclMap := make(map[string][]format.Operation)
	for _, acl := range acls {
		aclMap[acl.Principal] = append(aclMap[acl.Principal], acl.Operation)
	}
	acl := TopicACL{
		Consumer: []TopicUser{},
		Producer: []TopicUser{},
	}
	for p, a := range aclMap {
		if isConsumer(a) {
			acl.Consumer = append(acl.Consumer, TopicUser{
				Principal: p,
				Present:   true,
			})
		}
		if isProducer(a) {
			acl.Producer = append(acl.Producer, TopicUser{
				Principal: p,
				Present:   true,
			})
		}
	}

	return acl, nil
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
