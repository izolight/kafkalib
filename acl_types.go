package kafkalib

import (
	"bytes"
	"fmt"
	"text/tabwriter"
)

// ACLsByPrincipal contains all ACLs per Principal
type ACLsByPrincipal map[string]ACLs

func (a ACLsByPrincipal) MarshalText() ([]byte, error) {
	buf := bytes.Buffer{}
	w := tabwriter.NewWriter(&buf, 2, 8, 1, '\t', 0)
	_, err := w.Write([]byte(`Principal\tResourceType\tResourceName\tHost\tOperation\tPermission\n`))
	if err != nil {
		return nil, err
	}

	for p := range a {
		_, err = fmt.Fprintf(w, "%s\t\t\t\t\t\n", p)
		if err != nil {
			return nil, err
		}
		for _, acl := range a[p] {
			_, err = fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n",
				"", acl.ResourceType, acl.ResourceName, acl.Host, acl.Operation, acl.PermissionType)
			if err != nil {
				return nil, err
			}
		}
	}

	err = w.Flush()
	if err != nil {
		return nil, err
	}
	text := buf.Bytes()
	return text, nil
}

// ACLsByPrincipalAndResource contains all ACLs per Principal and Resource
type ACLsByPrincipalAndResource map[string]map[Resource]ACLs

// ACLsByResource contains all ACLs by Resource
type ACLsByResource map[Resource]ACLs

// ACLsByResourceAndPrincipal contains all ACLs by Resource and Principal
type ACLsByResourceAndPrincipal map[Resource]map[string]ACLs

// ACL holds all ACL information
// Kafka ACLs are defined in the general format of "Principal P is [Allowed/Denied] Operation O From Host H
// On Resource R matching ResourcePattern RP".
type ACL struct {
	Principal      string `json:"principal"`
	PermissionType string `json:"permission_type"`
	Operation      string `json:"operation"`
	Host           string `json:"host"`
	Resource
}

func (a ACL) String() string {
	return fmt.Sprintf("Principal %s is %s Operation %s from Host %s on Resource %s",
		a.Principal, a.PermissionType, a.Operation, a.Host, a.Resource)
}

func (a ACL) MarshalText() ([]byte, error) {
	buf := bytes.Buffer{}
	w := tabwriter.NewWriter(&buf, 2, 8, 1, '\t', 0)
	_, err := w.Write([]byte(`ResourceType\tResourceName\tPrincipal\tHost\tOperation\tPermission\n`))
	if err != nil {
		return nil, err
	}
	_, err = fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n",
		a.ResourceType, a.ResourceName, a.Principal, a.Host, a.Operation, a.PermissionType)
	if err != nil {
		return nil, err
	}
	err = w.Flush()
	if err != nil {
		return nil, err
	}
	text := buf.Bytes()
	return text, nil
}

type ACLs []ACL

func (a ACLs) MarshalText() ([]byte, error) {
	buf := bytes.Buffer{}
	w := tabwriter.NewWriter(&buf, 2, 8, 1, '\t', 0)
	_, err := w.Write([]byte(`ResourceType\tResourceName\tPrincipal\tHost\tOperation\tPermission\n`))
	if err != nil {
		return nil, err
	}
	for i := range a {
		_, err = fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n",
			a[i].ResourceType, a[i].ResourceName, a[i].Principal, a[i].Host, a[i].Operation, a[i].PermissionType)
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

// Resource holds the resource name and type
type Resource struct {
	ResourceName string `json:"resource_name"`
	ResourceType string `json:"resource_type"`
}

func (r Resource) String() string {
	return fmt.Sprintf("%s/%s", r.ResourceType, r.ResourceName)
}


