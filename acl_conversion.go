package kafkalib

import (
	"errors"
	"github.com/Shopify/sarama"
)

// ACLsByPrincipal contains all ACLs per Principal
type ACLsByPrincipal map[string][]ACL

// ACLsByPrincipalAndResource contains all ACLs per Principal and Resource
type ACLsByPrincipalAndResource map[string]map[Resource]ACLs

// ACLsByResource contains all ACLs by Resource
type ACLsByResource map[Resource]ACLs

// ACLsByResourceAndPrincipal contains all ACLs by Resource and Principal
type ACLsByResourceAndPrincipal map[Resource]map[string]ACLs

var (
	PermissionNotFound   = errors.New("Permission not found")
	OperationNotFound    = errors.New("Operation not found")
	ResourceTypeNotFound = errors.New("ResourceType not found")
)

// ACL holds all ACL information
// Kafka ACLs are defined in the general format of "Principal P is [Allowed/Denied] Operation O From Host H
// On Resource R matching ResourcePattern RP".
type ACL struct {
	Principal      string `json:"principal"`
	PermissionType string `json:"permission_type"`
	Operation      string `json:"operation"`
	Host           string `json:"host"`
	Resource       `json:"resource"`
}

type ACLs []ACL

// Resource holds the resource name and type
type Resource struct {
	ResourceName string `json:"resource_name"`
	ResourceType string `json:"resource_type"`
}

// MarshalSaramaRACL converts an ACL to sarama.ResourceAcl
func (a ACL) MarshalSaramaRACL() (*sarama.ResourceAcls, error) {
	acl, err := a.MarshalSaramaACL()
	if err != nil {
		return nil, err
	}

	resourceType, ok := resourceTypeToID[a.ResourceType]
	if !ok {
		return nil, ResourceTypeNotFound
	}

	return &sarama.ResourceAcls{
		Resource: sarama.Resource{
			ResourceName: a.ResourceName,
			ResourceType: resourceType,
		},
		Acls: []*sarama.Acl{acl},
	}, nil
}

// MarshalSaramaACL converts an ACL to sarama.Acl
func (a ACL) MarshalSaramaACL() (*sarama.Acl, error) {
	acl := &sarama.Acl{
		Principal: a.Principal,
		Host:      a.Host,
	}

	permissionType, ok := permissionToID[a.PermissionType]
	if !ok {
		return nil, PermissionNotFound
	}
	acl.PermissionType = permissionType

	operation, ok := operationToID[a.Operation]
	if !ok {
		return nil, OperationNotFound
	}
	acl.Operation = operation

	return acl, nil
}

// MarshalSarama converts a list of ACL(with same resource) to sarama.ResourceAcls
func (a ACLs) MarshalSaramaPerResource() (*sarama.ResourceAcls, error) {
	resourceType, ok := resourceTypeToID[a[0].ResourceType]
	if !ok {
		return nil, ResourceTypeNotFound
	}

	rACLs := &sarama.ResourceAcls{
		Resource: sarama.Resource{
			ResourceName: a[0].ResourceName,
			ResourceType: resourceType,
		},
		Acls: make([]*sarama.Acl, len(a)),
	}
	for i := range a {
		if a[i].ResourceName != rACLs.ResourceName {
			return nil, errors.New("ResourceMismatch")
		}
		resourceType, ok := resourceTypeToID[a[0].ResourceType]
		if !ok {
			return nil, ResourceTypeNotFound
		}
		if resourceType != rACLs.ResourceType {
			return nil, errors.New("ResourceMismatch")
		}
		acl, err := a[i].MarshalSaramaACL()
		if err != nil {
			return nil, err
		}
		rACLs.Acls[i] = acl
	}
	return rACLs, nil
}

// MarshalSarama converts ACLsByPrincipal to a list of sarama.ResourceAcls
func (a ACLsByPrincipal) MarshalSarama() ([]*sarama.ResourceAcls, error) {
	// convert to ACLsByResource for more efficiency
	abr := make(ACLsByResource)
	for p := range a {
		for _, acl := range a[p] {
			abr[acl.Resource] = append(abr[acl.Resource], acl)
		}
	}

	return abr.MarshalSarama()
}

// MarshalSarama converts ACLsByPrincipalAndResource to a list of sarama.ResourceAcls
func (a ACLsByPrincipalAndResource) MarshalSarama() ([]*sarama.ResourceAcls, error) {
	// convert to ACLsByResource for more efficiency
	abr := make(ACLsByResource)
	for p := range a {
		for r := range a[p] {
			abr[r] = append(abr[r], a[p][r]...)
		}
	}

	return abr.MarshalSarama()
}

// MarshalSarama converts ACLsByResource to a list of sarama.ResourceAcls
func (a ACLsByResource) MarshalSarama() ([]*sarama.ResourceAcls, error) {
	rACLs := make([]*sarama.ResourceAcls, len(a))
	i := 0
	for r := range a {
		acls, err := a[r].MarshalSaramaPerResource()
		if err != nil {
			return nil, err
		}
		rACLs[i] = acls
		i++
	}
	return rACLs, nil
}

// MarshalSarama converts ACLsByResourceAndPrincipal to a list of sarama.ResourceAcls
func (a ACLsByResourceAndPrincipal) MarshalSarama() ([]*sarama.ResourceAcls, error) {
	rACLs := make([]*sarama.ResourceAcls, len(a))
	i := 0
	for r := range a {
		resourceType, ok := resourceTypeToID[r.ResourceType]
		if !ok {
			return nil, ResourceTypeNotFound
		}
		apr := &sarama.ResourceAcls{
			Resource: sarama.Resource{
				ResourceType: resourceType,
				ResourceName: r.ResourceName,
			},
		}
		for p := range a[r] {
			acls, err := a[r][p].MarshalSaramaPerResource()
			if err != nil {
				return nil, err
			}
			apr.Acls = append(apr.Acls, acls.Acls...)
		}
		rACLs[i] = apr
		i++
	}
	return rACLs, nil
}

func (a *ACLs) UnmarshalSarama(rACLs *sarama.ResourceAcls) error {
	acls := make(ACLs, len(rACLs.Acls))
	resourceType, ok := resourceTypeToString[rACLs.ResourceType]
	if !ok {
		return ResourceTypeNotFound
	}
	i := 0
	for _, ac := range rACLs.Acls {
		acl := &ACL{
			Resource:Resource{
				ResourceType: resourceType,
				ResourceName: rACLs.ResourceName,
			},
		}
		err := acl.UnmarshalSarama(ac)
		if err != nil {
			return err
		}
		acls[i] = *acl
		i++
	}

	return nil
}

func (a *ACL) UnmarshalSarama(acl *sarama.Acl) error {
	a.Principal = acl.Principal
	a.Host = acl.Host
	operation, ok := operationToString[acl.Operation]
	if !ok {
		return OperationNotFound
	}
	a.Operation = operation

	permissionType, ok := permissionToString[acl.PermissionType]
	if !ok {
		return PermissionNotFound
	}
	a.PermissionType = permissionType

	return nil
}

// UnmarshalSarama converts a list of sarama.ResourceAcls to ACLsByResource
func (a ACLsByResource) UnmarshalSarama(rACLs []*sarama.ResourceAcls) error {
	for _, rACL := range rACLs {
		resourceType, ok := resourceTypeToString[rACL.ResourceType]
		if !ok {
			return ResourceTypeNotFound
		}
		resource := Resource{
			ResourceType: resourceType,
			ResourceName: rACL.ResourceName,
		}
		acls := &ACLs{}
		err := acls.UnmarshalSarama(rACL)
		if err != nil {
			return err
		}
		a[resource] = append(a[resource], *acls...)
	}
	return nil
}

// UnmarshalSarama converts a list of sarama.ResourceAcls to ACLsByResourceAndPrincipal
func (a ACLsByResourceAndPrincipal) UnmarshalSarama(rACLs []*sarama.ResourceAcls) error {
	for _, rACL := range rACLs {
		resourceType, ok := resourceTypeToString[rACL.ResourceType]
		if !ok {
			return ResourceTypeNotFound
		}
		resource := Resource{
			ResourceType: resourceType,
			ResourceName: rACL.ResourceName,
		}
		for _, ac := range rACL.Acls {
			acl := ACL{
				Resource: resource,
			}
			err := acl.UnmarshalSarama(ac)
			if err != nil {
				return err
			}
			a[resource][acl.Principal] = append(a[resource][acl.Principal], acl)
		}
	}

	return nil
}

func (a ACLsByPrincipal) UnmarshalSarama(rACLs []*sarama.ResourceAcls) error {
	for _, rACL := range rACLs {
		resourceType, ok := resourceTypeToString[rACL.ResourceType]
		if !ok {
			return ResourceTypeNotFound
		}
		resource := Resource{
			ResourceType: resourceType,
			ResourceName: rACL.ResourceName,
		}
		for _, ac := range rACL.Acls {
			acl := ACL{
				Resource: resource,
			}
			err := acl.UnmarshalSarama(ac)
			if err != nil {
				return err
			}
			a[acl.Principal] = append(a[acl.Principal], acl)
		}
	}

	return nil
}

func (a ACLsByPrincipalAndResource) UnmarshalSarama(rACLs []*sarama.ResourceAcls) error {
	for _, rACL := range rACLs {
		resourceType, ok := resourceTypeToString[rACL.ResourceType]
		if !ok {
			return ResourceTypeNotFound
		}
		resource := Resource{
			ResourceType: resourceType,
			ResourceName: rACL.ResourceName,
		}
		for _, ac := range rACL.Acls {
			acl := ACL{
				Resource: resource,
			}
			err := acl.UnmarshalSarama(ac)
			if err != nil {
				return err
			}
			a[acl.Principal][resource] = append(a[acl.Principal][resource], acl)
		}
	}

	return nil
}