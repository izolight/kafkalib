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

type ConsumerGroups []ConsumerGroup

func (c ConsumerGroups) MarshalText() ([]byte, error) {
	buf := bytes.Buffer{}
	w := tabwriter.NewWriter(&buf, 2, 8, 1, '\t', 0)
	_, err := w.Write([]byte(`ConsumerGroup\tACL\n`))
	if err != nil {
		return nil, err
	}
	for _, group := range c {
		_, err = fmt.Fprintf(w, "%s\t\n", group.Name)
		if err != nil {
			return nil, err
		}
		for _, acl := range group.ACLs {
			_, err = fmt.Fprintf(w, "%s\t%s\n",
				"", acl)
			if err != nil {
				return nil, err
			}
		}
	}
	err = w.Flush()
	if err != nil {
		return nil, err
	}
	text := buf.Bytes()
	return text, nil
}

func (c Conn) GetAllConsumerGroups() (ConsumerGroups, error) {
	groups, err := c.AdminClient.ListConsumerGroups()
	if err != nil {
		return nil, fmt.Errorf("Error getting all Consumergroups: %s", err)
	}
	g := ConsumerGroups{}
	for group := range groups {
		g = append(g, ConsumerGroup{
			Name: group,
		})
	}
	return g, nil
}

func (c Conn) DeleteConsumerGroup(name string) error {
	panic("not implemented")
}
