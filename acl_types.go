package kafkalib

// ACLsByPrincipal contains all ACLs per Principal
type ACLsByPrincipal map[string][]ACL

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

type ACLs []ACL

// Resource holds the resource name and type
type Resource struct {
	ResourceName string `json:"resource_name"`
	ResourceType string `json:"resource_type"`
}

