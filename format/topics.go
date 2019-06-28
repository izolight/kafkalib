package format

import (
	"encoding/json"
	"fmt"
	"sort"
	"text/tabwriter"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
)

// Topics holds the topic metadata received
type Topics map[string]sarama.TopicDetail

// Topic is a type alias
type Topic sarama.TopicDetail

// TopicDiff holds the difference between two topics
type TopicDiff map[string]map[string]interface{}

// Diff returns the difference between two topics and a boolean indicating if they diff
func (t *Topic) Diff(other *Topic) (TopicDiff, bool) {
	equal := true
	diff := TopicDiff{}
	if t.ReplicationFactor != other.ReplicationFactor {
		equal = false
		diff["replicas"]["left"] = t.ReplicationFactor
		diff["replicas"]["right"] = other.ReplicationFactor
	}
	if t.NumPartitions != other.NumPartitions {
		equal = false
		diff["partitions"]["left"] = t.NumPartitions
		diff["partitions"]["right"] = other.NumPartitions
	}
	if len(t.ConfigEntries) != len(other.ConfigEntries) {
		equal = false
		diff["config"]["left"] = t.ConfigEntries
		diff["config"]["right"] = t.ConfigEntries
	}
	for c := range t.ConfigEntries {
		if t.ConfigEntries[c] != other.ConfigEntries[c] {
			equal = false
		}
	}
	return diff, equal
}

// FormatText outputs the topic overview tab separated
func (topics Topics) FormatText(config Config) error {
	w := new(tabwriter.Writer)
	w.Init(config.Output, 0, 8, 0, '\t', 0)
	_, err := fmt.Fprintln(w, "Topic\tPartitions\tReplicationFactor")
	if err != nil {
		return err
	}
	for _, k := range config.TopicSort {
		_, err := fmt.Fprintf(w, "%s\t%d\t%d\n", k, topics[k].NumPartitions, topics[k].ReplicationFactor)
		if err != nil {
			return err
		}
	}
	err = w.Flush()
	if err != nil {
		return err
	}
	return nil
}

// FormatJSON implements the Formatter interface for Topics
func (topics Topics) FormatJSON(config Config) error {
	enc := json.NewEncoder(config.Output)
	if err := enc.Encode(topics); err != nil {
		return err
	}
	return nil
}

// FormatWide provides all output that is cut from normal FormatText
func (topics Topics) FormatWide(config Config) {
	w := new(tabwriter.Writer)
	w.Init(config.Output, 0, 8, 0, '\t', 0)
	_, err := fmt.Fprintln(w, "Topic\tPartitions\tReplicationfactor\tConfig Entries\tReplica Assignment")
	if err != nil {
		log.Fatal(err)
	}
	for _, k := range config.TopicSort {
		replicaAssignment, err := json.Marshal(topics[k].ReplicaAssignment)
		if err != nil {
			log.Fatal(err)
		}
		configEntries, err := json.Marshal(topics[k].ConfigEntries)
		if err != nil {
			log.Fatal(err)
		}
		_, err = fmt.Fprintf(w, "%s\t%d\t%d\t%s\t%s\n", k, topics[k].NumPartitions, topics[k].ReplicationFactor, configEntries, replicaAssignment)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = w.Flush()
	if err != nil {
		log.Fatal(err)
	}
}

// Sort returns a slice with the sorted keys for the topics map
func (topics Topics) Sort() []string {
	sorted := make([]string, 0, len(topics))
	for k := range topics {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)
	return sorted
}
