package entities

import (
	"encoding/json"
	"github.com/Shopify/sarama"
)

// ACL holds the acl basics independent of user or resource(type)
type ACL struct {
	Host           string `json:"host"`
	Operation      string `json:"operation"`
	PermissionType string `json:"permission_type"`
}

// ACLByPrincipal adds the resource and type to the acl
type ACLByPrincipal struct {
	ACLs         []ACL
	Resource     string `json:"resource"`
	ResourceType string `json:"resource-type"`
}

// ACLByResource adds the principal to the acl
type ACLByResource struct {
	Principal string `json:"principal"`
	ACLs      []ACL
}

// ACLsByPrincipal holds all acls with the principal as key
type ACLsByPrincipal map[string][]ACLByPrincipal

// aclsByPrincipal is used for json (un)marshalling of ACLsByPrincipal
type aclsByPrincipal struct {
	Principal string           `json:"principal"`
	Resources []ACLByPrincipal `json:"resources"`
}

// MarshalSarama converts ACLsByPrincipal to []sarama.ResourceAcls
func (a ACLsByPrincipal) MarshalSarama() ([]sarama.ResourceAcls, error) {
	abrt, err := a.marshalACLsByResourceType()
	if err != nil {
		return nil, err
	}
	return abrt.MarshalSarama()
}

// MarshalSarama converts ACLsByResourceType to []sarama.ResourceAcls
func (a ACLsByResourceType) MarshalSarama() ([]sarama.ResourceAcls, error) {
	out := []sarama.ResourceAcls{}
	for rt, abrts := range a {
		for r, abrs := range abrts {
			res := sarama.Resource{
				ResourceType: resourceTypeToID[rt],
				ResourceName: r,
			}
			acls := []*sarama.Acl{}
			for _, abr := range abrs {
				for _, acl := range abr.ACLs {
					acls = append(acls, &sarama.Acl{
						Principal:      abr.Principal,
						Host:           acl.Host,
						Operation:      operationToID[acl.Operation],
						PermissionType: permissionToID[acl.PermissionType],
					})
				}
			}
			out = append(out, sarama.ResourceAcls{
				Resource: res,
				Acls:     acls,
			})
		}
	}
	return out, nil
}

func (a ACLsByPrincipal) marshalACLsByResourceType() (ACLsByResourceType, error) {
	out := ACLsByResourceType{}
	for principal, abps := range a {
		for _, abp := range abps {
			if _, ok := out[abp.ResourceType]; !ok {
				out[abp.ResourceType] = ACLsByResource{}
			}
			if _, ok := out[abp.ResourceType][abp.Resource]; !ok {
				out[abp.ResourceType][abp.Resource] = []ACLByResource{}
			}
			out[abp.ResourceType][abp.Resource] = append(out[abp.ResourceType][abp.Resource], ACLByResource{
				Principal: principal,
				ACLs:      abp.ACLs,
			})
		}
	}
	return out, nil
}

// UnmarshalSarama converts []sarama.ResourceAcls to ACLsByPrincipal
func UnmarshalSarama(rAcls []sarama.ResourceAcls) (ACLsByPrincipal, error) {
	out := ACLsByPrincipal{}
	for _, rACL := range rAcls {
		for _, acl := range rACL.Acls {
			if _, ok := out[acl.Principal]; !ok {
				out[acl.Principal] = []ACLByPrincipal{
					{
						ResourceType: resourceTypeToString[rACL.ResourceType],
						Resource:     rACL.ResourceName,
					},
				}
			}
			found := false
			for i, abp := range out[acl.Principal] {
				if abp.ResourceType == resourceTypeToString[rACL.ResourceType] && abp.Resource == rACL.ResourceName {
					out[acl.Principal][i].ACLs = append(out[acl.Principal][i].ACLs, ACL{
						Operation:      operationToString[acl.Operation],
						PermissionType: permissionToString[acl.PermissionType],
						Host:           acl.Host,
					})
					found = true
					break
				}
			}
			if !found {
				out[acl.Principal] = append(out[acl.Principal], ACLByPrincipal{
					ResourceType: resourceTypeToString[rACL.ResourceType],
					Resource:     rACL.ResourceName,
					ACLs: []ACL{
						{
							Operation:      operationToString[acl.Operation],
							PermissionType: permissionToString[acl.PermissionType],
							Host:           acl.Host,
						},
					},
				})
			}
		}
	}
	return out, nil
}

// MarshalJSON ...
func (a ACLsByPrincipal) MarshalJSON() ([]byte, error) {
	out := []aclsByPrincipal{}
	for k, v := range a {
		out = append(out, aclsByPrincipal{Principal: k, Resources: v})
	}
	return json.Marshal(out)
}

// UnmarshalJSON ...
func (a ACLsByPrincipal) UnmarshalJSON(b []byte) error {
	tmp := []aclsByPrincipal{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}
	if a == nil {
		a = ACLsByPrincipal{}
	}
	for _, t := range tmp {
		a[t.Principal] = t.Resources
	}
	return nil
}

// ACLsByResource holds all acls with the resource name as key
type ACLsByResource map[string][]ACLByResource

// aclByResourceType is used for json (un)marshalling of ACLsByResource
type aclsByResource struct {
	Resource   string          `json:"resource"`
	Principals []ACLByResource `json:"principals"`
}

// MarshalJSON ...
func (a ACLsByResource) MarshalJSON() ([]byte, error) {
	out := []aclsByResource{}
	for k, v := range a {
		out = append(out, aclsByResource{Resource: k, Principals: v})
	}
	return json.Marshal(out)
}

// UnmarshalJSON ...
func (a ACLsByResource) UnmarshalJSON(b []byte) error {
	tmp := []aclsByResource{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}
	if a == nil {
		a = ACLsByResource{}
	}
	for _, t := range tmp {
		a[t.Resource] = t.Principals
	}
	return nil
}

// ACLsByResourceType holds the ACLsByResource with the type as key
type ACLsByResourceType map[string]ACLsByResource

// aclsByResourceType is used for json (un)marshalling of ACLsByResourceType
type aclsByResourceType struct {
	ResourceType string         `json:"resource-type"`
	Resources    ACLsByResource `json:"resources"`
}

// MarshalJSON ...
func (a ACLsByResourceType) MarshalJSON() ([]byte, error) {
	out := []aclsByResourceType{}
	for k, v := range a {
		out = append(out, aclsByResourceType{ResourceType: k, Resources: v})
	}
	return json.Marshal(out)
}

// UnmarshalJSON ...
func (a ACLsByResourceType) UnmarshalJSON(b []byte) error {
	tmp := []aclsByResourceType{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}
	if a == nil {
		a = ACLsByResourceType{}
	}
	for _, t := range tmp {
		a[t.ResourceType] = t.Resources
	}
	return nil
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

var operationToString = map[sarama.AclOperation]string{
	sarama.AclOperationUnknown:         "Unknown",
	sarama.AclOperationAny:             "Any",
	sarama.AclOperationAll:             "All",
	sarama.AclOperationRead:            "Read",
	sarama.AclOperationWrite:           "Write",
	sarama.AclOperationCreate:          "Create",
	sarama.AclOperationDelete:          "Delete",
	sarama.AclOperationAlter:           "Alter",
	sarama.AclOperationDescribe:        "Describe",
	sarama.AclOperationClusterAction:   "ClusterAction",
	sarama.AclOperationDescribeConfigs: "DescribeConfigs",
	sarama.AclOperationAlterConfigs:    "AlterConfigs",
	sarama.AclOperationIdempotentWrite: "IdempotentWrite",
}

var permissionToString = map[sarama.AclPermissionType]string{
	sarama.AclPermissionUnknown: "Unknown",
	sarama.AclPermissionAny:     "Any",
	sarama.AclPermissionDeny:    "Deny",
	sarama.AclPermissionAllow:   "Allow",
}

var resourceTypeToID = map[string]sarama.AclResourceType{
	"Unknown":         sarama.AclResourceUnknown,
	"Any":             sarama.AclResourceAny,
	"Topic":           sarama.AclResourceTopic,
	"Group":           sarama.AclResourceGroup,
	"Cluster":         sarama.AclResourceCluster,
	"TransactionalID": sarama.AclResourceTransactionalID,
}

var permissionToID = map[string]sarama.AclPermissionType{
	"Unknown": sarama.AclPermissionUnknown,
	"Any":     sarama.AclPermissionAny,
	"Deny":    sarama.AclPermissionDeny,
	"Allow":   sarama.AclPermissionAllow,
}

// operationToID maps a string to a acl operation
var operationToID = map[string]sarama.AclOperation{
	"Unknown":         sarama.AclOperationUnknown,
	"Any":             sarama.AclOperationAny,
	"All":             sarama.AclOperationAll,
	"Read":            sarama.AclOperationRead,
	"Write":           sarama.AclOperationWrite,
	"Create":          sarama.AclOperationCreate,
	"Delete":          sarama.AclOperationDelete,
	"Alter":           sarama.AclOperationAlter,
	"Describe":        sarama.AclOperationDescribe,
	"ClusterAction":   sarama.AclOperationClusterAction,
	"DescribeConfigs": sarama.AclOperationDescribeConfigs,
	"AlterConfigs":    sarama.AclOperationAlterConfigs,
	"IdempotentWrite": sarama.AclOperationIdempotentWrite,
}
