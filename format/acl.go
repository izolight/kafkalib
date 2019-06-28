package format

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"regexp"
	"text/tabwriter"
)

// ResourceType is a type alias
type ResourceType sarama.AclResourceType

// Permission is a type alias
type Permission sarama.AclPermissionType

// Operation is a type alias
type Operation sarama.AclOperation

// Resource contains the type and name of a resource
type Resource struct {
	Name string       `json:"name"`
	Type ResourceType `json:"type"`
}

// ACL is a container for working with acls
type ACL struct {
	Principal  string     `json:"principal"`
	Operation  Operation  `json:"operation"`
	Permission Permission `json:"permission"`
	Resource   Resource   `json:"resource"`
}

// ACLs is a container for grouping acls
type ACLs struct {
	ByResource  map[Resource][]ACL `json:"byResource"`
	ByPrincipal map[string][]ACL   `json:"byPrincipal"`
}

// maps for string conversion
var resourceTypeToString = map[sarama.AclResourceType]string{
	sarama.AclResourceUnknown:         "Unknown",
	sarama.AclResourceAny:             "Any",
	sarama.AclResourceTopic:           "Topic",
	sarama.AclResourceGroup:           "Group",
	sarama.AclResourceCluster:         "Cluster",
	sarama.AclResourceTransactionalID: "TransactionalID",
}

var resourceTypeToID = map[string]sarama.AclResourceType{
	"Unknown":         sarama.AclResourceUnknown,
	"Any":             sarama.AclResourceAny,
	"Topic":           sarama.AclResourceTopic,
	"Group":           sarama.AclResourceGroup,
	"Cluster":         sarama.AclResourceCluster,
	"TransactionalID": sarama.AclResourceTransactionalID,
}

var permissionToString = map[sarama.AclPermissionType]string{
	sarama.AclPermissionUnknown: "Unknown",
	sarama.AclPermissionAny:     "Any",
	sarama.AclPermissionDeny:    "Deny",
	sarama.AclPermissionAllow:   "Allow",
}

var permissionToID = map[string]sarama.AclPermissionType{
	"Unknown": sarama.AclPermissionUnknown,
	"Any":     sarama.AclPermissionAny,
	"Deny":    sarama.AclPermissionDeny,
	"Allow":   sarama.AclPermissionAllow,
}

var operationToString = map[sarama.AclOperation]string{
	sarama.AclOperationUnknown:         "Unknown",
	sarama.AclOperationAny:             "Any",
	sarama.AclOperationAll:             "All",
	sarama.AclOperationRead:            "Read",
	sarama.AclOperationWrite:           "Write",
	sarama.AclOperationCreate:          "CreateTopic",
	sarama.AclOperationDelete:          "DeleteTopic",
	sarama.AclOperationAlter:           "Alter",
	sarama.AclOperationDescribe:        "Describe",
	sarama.AclOperationClusterAction:   "Cluster Action",
	sarama.AclOperationDescribeConfigs: "Describe Configs",
	sarama.AclOperationAlterConfigs:    "Alter Configs",
	sarama.AclOperationIdempotentWrite: "Idempotent Write",
}

// OperationToID maps a string to a acl operation
var OperationToID = map[string]sarama.AclOperation{
	"Unknown":          sarama.AclOperationUnknown,
	"Any":              sarama.AclOperationAny,
	"All":              sarama.AclOperationAll,
	"Read":             sarama.AclOperationRead,
	"Write":            sarama.AclOperationWrite,
	"CreateTopic":      sarama.AclOperationCreate,
	"DeleteTopic":      sarama.AclOperationDelete,
	"Alter":            sarama.AclOperationAlter,
	"Describe":         sarama.AclOperationDescribe,
	"Cluster Action":   sarama.AclOperationClusterAction,
	"Describe Configs": sarama.AclOperationDescribeConfigs,
	"Alter Configs":    sarama.AclOperationAlterConfigs,
	"Idempotent Write": sarama.AclOperationIdempotentWrite,
}

// interface implementations

func (r ResourceType) String() string {
	return resourceTypeToString[sarama.AclResourceType(r)]
}

// MarshalJSON implements the Marshaler interface
func (r ResourceType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(resourceTypeToString[sarama.AclResourceType(r)])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON implements the Unmarshaler interface
func (r *ResourceType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*r = ResourceType(resourceTypeToID[j])
	return nil
}

// String implements the Stringer interface
func (p Permission) String() string {
	return permissionToString[sarama.AclPermissionType(p)]
}

// MarshalJSON implements the Marshaler interface
func (p Permission) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(permissionToString[sarama.AclPermissionType(p)])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON implements the Unmarshaler interface
func (p *Permission) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*p = Permission(permissionToID[j])
	return nil
}

// String implements the Stringer interface
func (o Operation) String() string {
	return operationToString[sarama.AclOperation(o)]
}

// MarshalJSON implements the Marshaler interface
func (o Operation) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(operationToString[sarama.AclOperation(o)])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON implements the Unmarshaler interface
func (o *Operation) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*o = Operation(OperationToID[j])
	return nil
}

// MarshalText implements the TextMarshaler interface
func (r Resource) MarshalText() (text []byte, err error) {
	return []byte(r.String()), nil
}

// String implements the Stringer interface
func (r Resource) String() string {
	return fmt.Sprintf("%s_%s", r.Type, r.Name)
}

// conversion

// FromResourceAcls converts sarama Acls to our ACLs
func FromResourceAcls(rs []sarama.ResourceAcls) *ACLs {
	acls := ACLs{
		ByResource:  make(map[Resource][]ACL),
		ByPrincipal: make(map[string][]ACL),
	}
	for _, r := range rs {
		for _, a := range r.Acls {
			acl := ACL{
				Principal:  a.Principal,
				Operation:  Operation(a.Operation),
				Permission: Permission(a.PermissionType),
				Resource: Resource{
					Name: r.ResourceName,
					Type: ResourceType(r.ResourceType),
				},
			}
			acls.ByPrincipal[acl.Principal] = append(acls.ByPrincipal[acl.Principal], acl)
			acls.ByResource[acl.Resource] = append(acls.ByResource[acl.Resource], acl)
		}
	}
	return &acls
}

// StringToACLFilter filters sarama Acls
// TODO: refactor to use our ACLs
func StringToACLFilter(filter string) (*sarama.AclFilter, error) {
	r, err := regexp.Compile(`(?P<type>topic|cluster|group|principal)/(?P<name>[\w.\-_:*]*)`)
	if err != nil {
		return nil, err
	}
	match := r.FindStringSubmatch(filter)
	resultMap := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i > 0 && i <= len(match) {
			resultMap[name] = match[i]
		}
	}
	af := sarama.AclFilter{}

	name := resultMap["name"]
	switch resultMap["type"] {
	case "topic":
		af.ResourceType = sarama.AclResourceTopic
		if len(name) > 0 {
			af.ResourceName = &name
		}
	case "cluster":
		af.ResourceType = sarama.AclResourceCluster
		if len(name) > 0 {
			af.ResourceName = &name
		}
	case "group":
		af.ResourceType = sarama.AclResourceGroup
		if len(name) > 0 {
			af.ResourceName = &name
		}
	case "principal":
		af.ResourceType = sarama.AclResourceAny
		if len(name) > 0 {
			af.Principal = &name
		}
	default:
		return nil, fmt.Errorf("Type %s is not any known acl type", resultMap["type"])
	}

	af.Operation = sarama.AclOperationAny
	af.PermissionType = sarama.AclPermissionAny

	return &af, nil
}

// formatting

func (acls ACLs) formatTextByResource(w *tabwriter.Writer) error {
	_, err := fmt.Fprintf(w, "ResourceType\tResourceName\tPrincipal\tOperation\tPermission\n")
	if err != nil {
		return err
	}
	for k, v := range acls.ByResource {
		_, err := fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", k.Type, k.Name, "", "", "")
		if err != nil {
			return err
		}
		for _, acl := range v {
			_, err := fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", "", "", acl.Principal, acl.Operation, acl.Permission)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (acls ACLs) formatTextByPrincipal(w *tabwriter.Writer) error {
	_, err := fmt.Fprintf(w, "Principal\tResourceType\tResourceName\tOperation\tPermission\n")
	if err != nil {
		return err
	}
	for k, v := range acls.ByPrincipal {
		_, err := fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", k, "", "", "", "")
		if err != nil {
			return err
		}
		for _, acl := range v {
			_, err := fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", "", acl.Resource.Type, acl.Resource.Name, acl.Operation, acl.Permission)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// FormatText implements the Formatter interface for ACLs
func (acls ACLs) FormatText(config Config) error {
	w := new(tabwriter.Writer)
	w.Init(config.Output, 2, 8, 1, '\t', 0)
	var err error
	switch config.ACLOrder {
	case "resource":
		err = acls.formatTextByResource(w)
	case "principal":
		err = acls.formatTextByPrincipal(w)
	default:
		err = acls.formatTextByResource(w)
	}
	if err != nil {
		return err
	}
	err = w.Flush()
	if err != nil {
		return err
	}
	return nil
}

// FormatJSON implements the Formatter interface for ACLs
func (acls ACLs) FormatJSON(config Config) error {
	enc := json.NewEncoder(config.Output)
	switch config.ACLOrder {
	case "resource":
		if err := enc.Encode(acls.ByResource); err != nil {
			return err
		}
	case "principal":
		if err := enc.Encode(acls.ByPrincipal); err != nil {
			return err
		}
	default:
		if err := enc.Encode(acls.ByResource); err != nil {
			return err
		}
	}
	return nil
}
