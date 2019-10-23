package kafkalib

import (
	"bytes"
	"fmt"
	"text/tabwriter"
)

// Topic represents the serialized Topic
type Topic struct {
	Name              string `json:"name"`
	Partitions        int    `json:"partitions"`
	ReplicationFactor int    `json:"replication_factor"`
	RetentionMs       int    `json:"retention_ms,omitempty"`
	ACLs              []ACLs `json:"acls,omitempty"`
}

// Topics is a type alias for a slice of topics
type Topics []Topic

// MarshalText prints the topic config in a tab separated view into a buffer
func (t Topics) MarshalText() ([]byte, error) {
	buf := bytes.Buffer{}
	w := tabwriter.NewWriter(&buf, 2, 8, 1, '\t', 0)
	_, err := w.Write([]byte(`Topic\tPartitions\tReplicationFactor\tRetention\n`))
	if err != nil {
		return nil, err
	}
	for _, topic := range t {
		_, err = fmt.Fprintf(w, "%s\t%d\t%d\t%d\t\n",
			topic.Name, topic.Partitions, topic.ReplicationFactor, topic.RetentionMs)
		if err != nil {
			return nil, err
		}
		for _, acl := range topic.ACLs {
			_, err = fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", "", "", "", "", acl)
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
