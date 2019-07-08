package kafkalib_test

import (
	"fmt"
	"github.com/Shopify/sarama"
)

type testClient struct {
	topics map[string]sarama.TopicDetail
	acls   []sarama.ResourceAcls
}

func (t *testClient) DeleteConsumerGroup(group string) error {
	panic("implement me")
}

// NewTestClient creates a client that is used in tests
func NewTestClient() sarama.ClusterAdmin {
	admin := &testClient{}
	admin.topics = map[string]sarama.TopicDetail{
		"simpleTopic": sarama.TopicDetail{
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
		"topicWithReplicas": sarama.TopicDetail{
			NumPartitions:     1,
			ReplicationFactor: 3,
		},
		"topicWithPartitions": sarama.TopicDetail{
			NumPartitions:     4,
			ReplicationFactor: 1,
		},
		"topicWithPartitionsAndReplicas": sarama.TopicDetail{
			NumPartitions:     4,
			ReplicationFactor: 3,
		},
	}
	admin.acls = []sarama.ResourceAcls{
		{
			Resource: sarama.Resource{ResourceType: sarama.AclResourceTopic, ResourceName: "test"},
			Acls: []*sarama.Acl{
				{Principal: "User:test", Host: "localhost", Operation: sarama.AclOperationAll, PermissionType: sarama.AclPermissionAllow},
				{Principal: "User:test2", Host: "localhost", Operation: sarama.AclOperationCreate, PermissionType: sarama.AclPermissionAllow},
			}},
	}
	return admin
}

func (t *testClient) CreateTopic(topic string, detail *sarama.TopicDetail, validateOnly bool) error {
	if _, ok := t.topics[topic]; ok {
		return fmt.Errorf("Topic %s already exists", topic)
	}
	t.topics[topic] = *detail
	return nil
}

func (t *testClient) ListTopics() (map[string]sarama.TopicDetail, error) {
	if t.topics != nil {
		return t.topics, nil
	}
	return nil, fmt.Errorf("Error getting topics")
}

func (t *testClient) DescribeTopics(topics []string) (metadata []*sarama.TopicMetadata, err error) {
	panic("not implemented")
}

func (t *testClient) DeleteTopic(topic string) error {
	if _, ok := t.topics[topic]; ok {
		delete(t.topics, topic)
		return nil
	}
	return fmt.Errorf("Topic %s not found", topic)
}

func (t *testClient) CreatePartitions(topic string, count int32, assignment [][]int32, validateOnly bool) error {
	panic("not implemented")
}

func (t *testClient) DeleteRecords(topic string, partitionOffsets map[int32]int64) error {
	panic("not implemented")
}

func (t *testClient) DescribeConfig(resource sarama.ConfigResource) ([]sarama.ConfigEntry, error) {
	panic("not implemented")
}

func (t *testClient) AlterConfig(resourceType sarama.ConfigResourceType, name string, entries map[string]*string, validateOnly bool) error {
	panic("not implemented")
}

func (t *testClient) CreateACL(resource sarama.Resource, acl sarama.Acl) error {
	if resource.ResourceType == sarama.AclResourceTopic {
		_, ok := t.topics[resource.ResourceName]
		if !ok {
			return fmt.Errorf("Can't create acl for non exisiting topic %s", resource.ResourceName)
		}
		for i, ta := range t.acls {
			if ta.ResourceName == resource.ResourceName {
				for _, a := range ta.Acls {
					if a.PermissionType == acl.PermissionType && a.Principal == acl.Principal && a.Operation == acl.Operation {
						return fmt.Errorf("Acl already exists")
					}
				}
				t.acls[i].Acls = append(t.acls[i].Acls, &acl)
				return nil
			}
		}
	}
	return nil
}

func (t *testClient) ListAcls(filter sarama.AclFilter) ([]sarama.ResourceAcls, error) {
	if t.acls != nil {
		return t.acls, nil
	}
	return nil, fmt.Errorf("Error getting acls")
}

func (t *testClient) DeleteACL(filter sarama.AclFilter, validateOnly bool) ([]sarama.MatchingAcl, error) {
	panic("not implemented")
}

func (t *testClient) ListConsumerGroups() (map[string]string, error) {
	panic("not implemented")
}

func (t *testClient) DescribeConsumerGroups(groups []string) ([]*sarama.GroupDescription, error) {
	panic("not implemented")
}

func (t *testClient) ListConsumerGroupOffsets(group string, topicPartitions map[string][]int32) (*sarama.OffsetFetchResponse, error) {
	panic("not implemented")
}

func (t *testClient) DescribeCluster() (brokers []*sarama.Broker, controllerID int32, err error) {
	panic("not implemented")
}

func (t *testClient) Close() error {
	t.topics = nil
	t.acls = nil
	return nil
}
