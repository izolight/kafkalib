package kafkalib

import (
	"bytes"
	"encoding/json"
)

// just for sanity tests
type _acls ACLs
func (a _acls) MarshalJSON() ([]byte, error) {
	return json.Marshal(ACLs(a))
}

// just for sanity tests
type _abp ACLsByPrincipal
func (a _abp) MarshalJSON() ([]byte, error) {
	return json.Marshal(ACLsByPrincipal(a))
}

func (a ACLsByResource) MarshalJSON() ([]byte, error) {
	length := len(a)
	buffer := bytes.NewBufferString(`{`)
	i := 1
	for r := range a {
		buffer.WriteString(`"` + r.ResourceType + `_` + r.ResourceName + `":`)
		b, err := json.Marshal(a[r])
		if err != nil {
			return nil, err
		}
		buffer.Write(b)
		if i != length {
			buffer.WriteString(`,`)
		}
		i++
	}
	buffer.WriteString(`}`)
	return buffer.Bytes(), nil
}