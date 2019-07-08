package kafkalib

import (
	"bytes"
	"fmt"
	"text/tabwriter"
)

// ConsumerGroup represents the serialized Consumergroup state
type ConsumerGroup struct {
	Name string `json:"name"`
	ACLs []ACL  `json:"acls"`
}

func (c ConsumerGroup) MarshalText() ([]byte, error) {
	buf := bytes.Buffer{}
	w := tabwriter.NewWriter(&buf, 2, 8, 1, '\t', 0)
	_, err := w.Write([]byte(`ConsumerGroup\tPrincipal\tHost\tOperation\tPermission\n`))
	if err != nil {
		return nil, err
	}
	_, err = fmt.Fprintf(w, "%s\t\t\t\t\n", c.Name)
	if err != nil {
		return nil, err
	}
	for _, acl := range c.ACLs {
		_, err = fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
			"", acl.Principal, acl.Host, acl.Operation, acl.PermissionType)
		if err != nil {
			return nil, err
		}
	}

	err = w.Flush()
	if err != nil {
		return nil, err
	}
	text := buf.Bytes()
	return text, nil
}