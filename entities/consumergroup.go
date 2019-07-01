package entities

// ConsumerGroup represents the serialized Consumergroup state
type ConsumerGroup struct {
	Name    string             `json:"name"`
	Present state              `json:"state"`
	ACL     []ConsumerGroupACL `json:"acl"`
}

// ConsumerGroupACL represents the serialized User on a Consumergroup
type ConsumerGroupACL struct {
	Principal   string   `json:"principal"`
	Present     state    `json:"state"`
	Permissions []string `json:"permission"`
}
