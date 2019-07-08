package kafkalib

import (
	"crypto/tls"
	"github.com/Shopify/sarama"
)

// Conn is a type alias
type Conn struct {
	AdminClient sarama.ClusterAdmin
	Client      sarama.Client
	Consumer    sarama.Consumer
}

// Config holds the config values for connecting to kafka
type Config struct {
	BrokerList            []string
	TLSEnabled            bool
	TLSInsecureSkipVerify bool
	User                  string
	Password              string
}

// NewAdminClient wraps the Admin creation of sarama
func NewAdminClient(config *Config) (sarama.ClusterAdmin, error) {
	cfg := sarama.NewConfig()
	cfg.ClientID = "kafkactl"
	cfg.Version = sarama.V2_0_0_0
	cfg.Net.TLS.Enable = config.TLSEnabled
	cfg.Net.TLS.Config = &tls.Config{
		InsecureSkipVerify: config.TLSInsecureSkipVerify,
	}
	if len(config.User) != 0 && len(config.Password) != 0 {
		cfg.Net.SASL.Enable = true
		cfg.Net.SASL.User = config.User
		cfg.Net.SASL.Password = config.Password
	}

	admin, err := sarama.NewClusterAdmin(config.BrokerList, cfg)
	if err != nil {
		return nil, err
	}
	return admin, err
}

// NewConsumer wraps the Consumer creation of sarama
func NewConsumer(config *Config) (sarama.Consumer, error) {
	cfg := sarama.NewConfig()
	cfg.ClientID = "kafkactl"
	cfg.Version = sarama.V2_0_0_0
	cfg.Net.TLS.Enable = config.TLSEnabled
	cfg.Net.TLS.Config = &tls.Config{
		InsecureSkipVerify: config.TLSInsecureSkipVerify,
	}
	if len(config.User) != 0 && len(config.Password) != 0 {
		cfg.Net.SASL.Enable = true
		cfg.Net.SASL.User = config.User
		cfg.Net.SASL.Password = config.Password
	}
	consumer, err := sarama.NewConsumer(config.BrokerList, cfg)
	if err != nil {
		return nil, err
	}
	return consumer, err
}

// NewClient wraps the Client creation of sarama
func NewClient(config *Config) (sarama.Client, error) {
	cfg := sarama.NewConfig()
	cfg.ClientID = "kafkactl"
	cfg.Version = sarama.V2_0_0_0
	cfg.Net.TLS.Enable = config.TLSEnabled
	cfg.Net.TLS.Config = &tls.Config{
		InsecureSkipVerify: config.TLSInsecureSkipVerify,
	}
	if len(config.User) != 0 && len(config.Password) != 0 {
		cfg.Net.SASL.Enable = true
		cfg.Net.SASL.User = config.User
		cfg.Net.SASL.Password = config.Password
	}
	client, err := sarama.NewClient(config.BrokerList, cfg)
	if err != nil {
		return nil, err
	}
	return client, err
}
