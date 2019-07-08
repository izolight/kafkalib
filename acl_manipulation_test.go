package kafkalib_test

import (
	"github.com/Shopify/sarama"
	"github.com/izolight/kafkalib/client_test"
	"testing"
)

func TestACL_Get(t *testing.T) {

}

func TestACL_GetAll(t *testing.T) {
	client := client_test.NewTestClient()
	testCases := []struct {
		success bool
	}{
		{true},
		{false},
	}
	for _, tc := range testCases {
		c := Conn{
			AdminClient: client,
		}
		filter := &sarama.AclFilter{
			ResourceType:   sarama.AclResourceAny,
			Operation:      sarama.AclOperationAny,
			PermissionType: sarama.AclPermissionAny,
		}
		if !tc.success {
			client.Close()
		}
		acls, err := c.GetACLs(filter)
		if err != nil && tc.success {
			t.Fatal(err)
		}
		if acls != nil && !tc.success {
			t.Fatal("Returned acls even though it shouldn't")
		}
	}
}

func TestACL_Create(t *testing.T) {
	client := client_test.NewTestClient()
	testCases := []struct {
		acl     sarama.AclCreation
		success bool
	}{
		{sarama.AclCreation{
			sarama.Resource{ResourceName: "simpleTopic", ResourceType: sarama.AclResourceTopic},
			sarama.Acl{Operation: sarama.AclOperationRead, Principal: "User:test", PermissionType: sarama.AclPermissionAllow},
		},
			true,
		},
		{sarama.AclCreation{
			sarama.Resource{ResourceName: "topicDoesNotExist", ResourceType: sarama.AclResourceTopic},
			sarama.Acl{Operation: sarama.AclOperationWrite, Principal: "User:test", PermissionType: sarama.AclPermissionAllow},
		},
			false,
		},
	}
	for _, tc := range testCases {
		c := Conn{
			AdminClient: client,
		}
		err := c.CreateACL(&tc.acl)
		if err != nil && tc.success {
			t.Fatal(err)
		}
		if err == nil && !tc.success {
			t.Fatal("Should return an error")
		}
	}
}

func TestACL_Delete(t *testing.T) {

}
