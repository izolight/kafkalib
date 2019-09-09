package kafkalib

import "github.com/Shopify/sarama"

func (t Topic) MarshallSarama() (*sarama.TopicDetail, error) {
	s := &sarama.TopicDetail{
		NumPartitions: int32(t.Partitions),
		ReplicationFactor: int16(t.ReplicationFactor),
	}

	return s, nil
}

func (t Topics) MarshallSarama() (map[string]sarama.TopicDetail, error) {
	s := make(map[string]sarama.TopicDetail)
	for _, topic := range t {
		s[topic.Name] = sarama.TopicDetail{
			NumPartitions: int32(topic.Partitions),
			ReplicationFactor: int16(topic.ReplicationFactor),
		}
	}
	return s, nil
}

func (t Topic) UnmarshallSarama(detail *sarama.TopicDetail) error {
	t.Partitions = int(detail.NumPartitions)
	t.ReplicationFactor = int(detail.ReplicationFactor)
	return nil
}

func (t Topics) UnmarshallSarama(details map[string]sarama.TopicDetail) error {
	for name, topic := range details {
		t = append(t, Topic{
			Name: name,
			Partitions: int(topic.NumPartitions),
			ReplicationFactor: int(topic.ReplicationFactor),
		})
	}
	return nil
}
