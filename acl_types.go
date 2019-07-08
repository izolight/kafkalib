package kafkalib

import "fmt"

// ACLsByPrincipal contains all ACLs per Principal
type ACLsByPrincipal map[string]ACLs

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

type ACLs []ACL

// Resource holds the resource name and type
type Resource struct {
	ResourceName string `json:"resource_name"`
	ResourceType string `json:"resource_type"`
}

func (r Resource) String() string {
	return fmt.Sprintf("%s/%s", r.ResourceType, r.ResourceName)
}
