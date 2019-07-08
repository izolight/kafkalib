package kafkalib

import (
	"fmt"
	"github.com/Shopify/sarama"
)

// GetACLs returns ACLs according to filter
func (c Conn) GetACLs(filter *sarama.AclFilter) (*format.ACLs, error) {
	var err error
	acls, err := c.AdminClient.ListAcls(*filter)
	if err != nil {
		return nil, fmt.Errorf("Error getting acls: %s", err)
	}
	return format.FromResourceAcls(acls), nil
}

// CreateACL creates a new acl
func (c Conn) CreateACL(acl *sarama.AclCreation) error {
	err := c.AdminClient.CreateACL(acl.Resource, acl.Acl)
	return err
}

// DeleteACL deletes a ACL according to filter
func (c Conn) DeleteACL(filter *sarama.AclFilter) ([]sarama.MatchingAcl, error) {
	acls, err := c.AdminClient.DeleteACL(*filter, false)
	return acls, err
}
