package format

import (
	"io"
)

// Formatter provides some methods for outputting the kafka metadata
type Formatter interface {
	FormatText(config Config) error
	FormatJSON(config Config) error
}

// Config contains options for the formatting
type Config struct {
	Output    io.Writer
	Format    string
	TopicSort []string
	ACLOrder  string
}

// Format is a wrapper around the different Format Methods
func Format(f Formatter, config Config) error {
	switch config.Format {
	case "text":
		return f.FormatText(config)
	case "json":
		return f.FormatJSON(config)
	default:
		return f.FormatText(config)
	}
}
