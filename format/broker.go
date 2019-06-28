package format

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"text/tabwriter"
)

// Brokers is a type alias
type Brokers []*sarama.Broker

// FormatText outputs the broker overview tab separated
func (brokers Brokers) FormatText(config Config) error {
	w := new(tabwriter.Writer)
	w.Init(config.Output, 0, 8, 0, '\t', 0)
	_, err := fmt.Fprintln(w, "Id\tAddress\tRack")
	if err != nil {
		return err
	}
	for _, k := range brokers {
		_, err := fmt.Fprintf(w, "%d\t%s\t%s\n", k.ID(), k.Addr(), k.Rack())
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

// FormatJSON outputs the broker overview tab separated
func (brokers Brokers) FormatJSON(config Config) error {
	enc := json.NewEncoder(config.Output)
	if err := enc.Encode(brokers); err != nil {
		return err
	}
	return nil
}
