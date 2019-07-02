package format_test

import (
	"bytes"
	"github.com/Shopify/sarama"
	"gitea.izolight.xyz/gabor/kafkalib/format"
	"strings"
	"testing"
)

func TestACLs_FormatJSON(t *testing.T) {
	acl := format.ACL{
		Principal:  "User:test",
		Operation:  format.Operation(sarama.AclOperationAll),
		Permission: format.Permission(sarama.AclPermissionAllow),
		Resource:   format.Resource{"test", format.ResourceType(sarama.AclResourceTopic)},
	}
	acls := format.ACLs{
		ByPrincipal: map[string][]format.ACL{
			"User:test": {acl},
		},
		ByResource: map[format.Resource][]format.ACL{
			acl.Resource: {acl},
		},
	}

	testCases := []struct {
		acls     format.ACLs
		order    string
		expected string
	}{
		{
			acls,
			"resource",
			`{"Topic_test":[{"principal":"User:test","operation":"All","permission":"Allow","resource":"Topic_test"}]}`,
		},
		{
			acls,
			"principal",
			`{"User:test":[{"principal":"User:test","operation":"All","permission":"Allow","resource":"Topic_test"}]}`,
		},
	}
	for _, tc := range testCases {
		output := new(bytes.Buffer)
		cfg := format.Config{
			Output:   output,
			Format:   "json",
			ACLOrder: tc.order,
		}
		format.Format(tc.acls, cfg)
		got := strings.TrimSuffix(output.String(), "\n")
		if got != tc.expected {
			t.Errorf("acls.FormatJSON():\nGot:\t%s\nWant:\t%s", got, tc.expected)
		}
	}
}

func TestACLs_FormatText(t *testing.T) {
	acl := format.ACL{
		Principal:  "User:test",
		Operation:  format.Operation(sarama.AclOperationAll),
		Permission: format.Permission(sarama.AclPermissionAllow),
		Resource:   format.Resource{"test", format.ResourceType(sarama.AclResourceTopic)},
	}
	acls := format.ACLs{
		ByPrincipal: map[string][]format.ACL{
			"User:test": {acl},
		},
		ByResource: map[format.Resource][]format.ACL{
			acl.Resource: {acl},
		},
	}

	testCases := []struct {
		acls     format.ACLs
		order    string
		expected string
	}{
		{
			acls,
			"resource",
			"ResourceType	ResourceName	Principal	Operation	Permission\n" +
				"Topic		test						\n" +
				"				User:test	All		Allow",
		},
		{
			acls,
			"principal",
			"Principal	ResourceType	ResourceName	Operation	Permission\n" +
				"User:test							\n" +
				"		Topic		test		All		Allow",
		},
	}
	for _, tc := range testCases {
		output := new(bytes.Buffer)
		cfg := format.Config{
			Output:   output,
			Format:   "text",
			ACLOrder: tc.order,
		}
		format.Format(tc.acls, cfg)
		got := strings.TrimSuffix(output.String(), "\n")
		if got != tc.expected {
			t.Errorf("acls.FormatText():\nGot:\t%s\nWant:\t%s", got, tc.expected)
		}
	}
}
