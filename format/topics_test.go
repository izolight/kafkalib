package format_test

import (
	"bytes"
	"github.com/Shopify/sarama"
	"github.com/izolight/kafkalib/format"
	"strings"
	"testing"
)

func TestTopics_FormatJSON(t *testing.T) {
	topics := format.Topics{
		"simpleTopic": sarama.TopicDetail{
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}
	testCases := []struct {
		topics   format.Topics
		expected string
	}{
		{
			topics,
			`{"simpleTopic":{"NumPartitions":1,"ReplicationFactor":1,"ReplicaAssignment":null,"ConfigEntries":null}}`,
		},
	}
	for _, tc := range testCases {
		output := new(bytes.Buffer)
		cfg := format.Config{
			Output:    output,
			Format:    "json",
			TopicSort: topics.Sort(),
		}
		format.Format(topics, cfg)
		got := strings.TrimSuffix(output.String(), "\n")
		if got != tc.expected {
			t.Errorf("topics.FormatJSON():\nGot:\t%s\nWant:\t%s", got, tc.expected)
		}
	}
}

func TestTopics_FormatText(t *testing.T) {
	topics := format.Topics{
		"simpleTopic": sarama.TopicDetail{
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}
	testCases := []struct {
		topics   format.Topics
		expected string
	}{
		{
			topics,
			"Topic\t\tPartitions\tReplicationFactor\n" +
				"simpleTopic\t1\t\t1",
		},
	}
	for _, tc := range testCases {
		output := new(bytes.Buffer)
		cfg := format.Config{
			Output:    output,
			Format:    "text",
			TopicSort: topics.Sort(),
		}
		format.Format(topics, cfg)
		got := strings.TrimSuffix(output.String(), "\n")
		if got != tc.expected {
			t.Errorf("topics.FormatText():\nGot:\t%s\nWant:\t%s", got, tc.expected)
		}
	}
}
