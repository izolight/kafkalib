package kafkalib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// just for sanity tests
type _acls ACLs

func (a _acls) UnmarshalJSON(b []byte) error {
	acls := ACLs(a)
	return json.Unmarshal(b, &acls)
}

func (a _acls) MarshalJSON() ([]byte, error) {
	return json.Marshal(ACLs(a))
}

// just for sanity tests
type _abp ACLsByPrincipal

func (a _abp) UnmarshalJSON(b []byte) error {
	abp := ACLsByPrincipal(a)
	return json.Unmarshal(b, &abp)
}

func (a _abp) MarshalJSON() ([]byte, error) {
	return json.Marshal(ACLsByPrincipal(a))
}

func (a ACLsByResource) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`{`)
	length := len(a)
	i := 1
	for r := range a {
		buffer.WriteString(`"` + r.ResourceType + `/` + r.ResourceName + `":`)
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

func (a ACLsByPrincipalAndResource) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`{`)
	principals := len(a)
	i := 1
	for p := range a {
		buffer.WriteString(`"` + p + `":[`)
		resources := len(a[p])
		j := 1
		for r := range a[p] {
			buffer.WriteString(`{"` + r.ResourceType + `/` + r.ResourceName + `":`)
			b, err := json.Marshal(a[p][r])
			if err != nil {
				return nil, err
			}
			buffer.Write(b)
			buffer.WriteString(`}`)
			if j != resources {
				buffer.WriteString(`,`)
			}
			j++
		}
		buffer.WriteString(`]`)
		if i != principals {
			buffer.WriteString(`,`)
		}
		i++
	}

	buffer.WriteString(`}`)
	return buffer.Bytes(), nil
}

func (a ACLsByResourceAndPrincipal) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`{`)
	resources := len(a)
	i := 1
	for r := range a {
		buffer.WriteString(`"` + r.ResourceType + `/` + r.ResourceName + `":[`)
		principals := len(a[r])
		j := 1
		for p := range a[r] {
			buffer.WriteString(`{"` + p + `":`)
			b, err := json.Marshal(a[r][p])
			if err != nil {
				return nil, err
			}
			buffer.Write(b)
			buffer.WriteString(`}`)
			if j != principals {
				buffer.WriteString(`,`)
			}
			j++
		}
		buffer.WriteString(`]`)
		if i != resources {
			buffer.WriteString(`,`)
		}
		i++
	}


	buffer.WriteString(`}`)
	return buffer.Bytes(), nil
}

func (a ACLsByResource) UnmarshalJSON(b []byte) error {
	tmp := map[string]ACLs{}
	err := json.Unmarshal(b, tmp)
	if err != nil {
		return err
	}
	for r := range tmp {
		parts := strings.Split(r, "/")
		if len(parts) != 2 {
			return fmt.Errorf("returned more than two parts when splitting at '/': %s", parts)
		}
		a[Resource{
			ResourceType: parts[0],
			ResourceName: parts[1],
		}] = tmp[r]
	}

	return nil
}

func (a ACLsByResourceAndPrincipal) UnmarshalJSON(b []byte) error {
	return nil
}