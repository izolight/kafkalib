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

type Topics []Topic

func (t Topic) MarshalText() ([]byte, error) {
	buf := bytes.Buffer{}
	w := tabwriter.NewWriter(&buf, 2, 8, 1, '\t', 0)
	_, err := w.Write([]byte(`Topic\tPartitions\tReplicationFactor\tRetention\tACLs\n`))
	if err != nil {
		return nil, err
	}
	_, err = fmt.Fprintf(w, "%s\t%d\t%d\t%d\t\n",
		t.Name, t.Partitions, t.ReplicationFactor, t.RetentionMs)
	if err != nil {
		return nil, err
	}
	for _, acl := range t.ACLs {
		_, err = fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", "", "", "", "", acl)
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