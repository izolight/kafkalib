package entities

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type state bool

func (s *state) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	if j == "present" {
		*s = true
		return nil
	} else if j == "absent" {
		*s = false
		return nil
	}
	return fmt.Errorf("%s is not a valid value for state", j)
}

func (s state) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	if s == true {
		buffer.WriteString("present")
	} else {
		buffer.WriteString("absent")
	}
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}
