package kafkalib

import "fmt"

func (c Conn) GetAllTopics() (Topics, error) {
	var err error
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