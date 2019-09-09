package kafkalib

import (
	"fmt"
	"regexp"
)

// GetAllTopics returns all known topics
func (c Conn) GetAllTopics() (Topics, error) {
	topics, err := c.AdminClient.ListTopics()
	if err != nil {
		return nil, fmt.Errorf("Error getting topics: %s", err)
	}
	t := Topics{}
	err = t.UnmarshallSarama(topics)
	if err != nil {
		return nil, fmt.Errorf("Error converting topics: %s", err)
	}
	return t, nil
}

// GetTopic returns the topic defined in the Name of the client

func (c Conn) GetTopic(filter string) (Topics, error) {
	all, err := c.AdminClient.ListTopics()
	if err != nil {
		return nil, fmt.Errorf("Error getting topics: %s", err)
	}
	t := Topics{}
	r, err := regexp.Compile(filter)
	for k, v := range all {
		if r.MatchString(k) {
			t = append(t, Topic{
				Name: k,
				Partitions: int(v.NumPartitions),
				ReplicationFactor: int(v.ReplicationFactor),
			})
		}
	}
	if len(t) > 0 {
		return t, nil
	}
	return nil, fmt.Errorf("Topic %s not found", filter)
}