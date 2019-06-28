package format

import (
	"encoding/json"
	"fmt"
	"text/tabwriter"
)

// ConsumerGroups is a type alias
type ConsumerGroups map[string]string

// FormatJSON implements the Formatter interface
func (cg ConsumerGroups) FormatJSON(config Config) error {
	enc := json.NewEncoder(config.Output)
	if err := enc.Encode(cg); err != nil {
		return err
	}
	return nil
}

// FormatText implements the FormatText interface
func (cg ConsumerGroups) FormatText(config Config) error {
	w := new(tabwriter.Writer)
	w.Init(config.Output, 0, 8, 0, '\t', 0)
	_, err := fmt.Fprintln(w, "Consumergroup\tConsumer")
	if err != nil {
		return err
	}
	for k, v := range cg {
		_, err := fmt.Fprintf(w, "%s\t%s\n", k, v)
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
