package kafkalib

// ConsumerGroup represents the serialized Consumergroup state
type ConsumerGroup struct {
	Name string `json:"name"`
	ACLs []ACL  `json:"acls"`
}
