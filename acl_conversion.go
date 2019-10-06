package kafkalib

import (
	"errors"
	"github.com/Shopify/sarama"
)

var (
	PermissionNotFound   = errors.New("Permission not found")
	OperationNotFound    = errors.New("Operation not found")
	ResourceTypeNotFound = errors.New("ResourceType not found")
)

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

	operation, ok := OperationToID[a.Operation]
	if !ok {
		return nil, OperationNotFound
	}
	acl.Operation = operation

	return acl, nil
}

// MarshalResourceAcls converts a list of ACL(with same resource) to sarama.ResourceAcls
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

// MarshalResourceAcls converts ACLsByPrincipal to a list of sarama.ResourceAcls
func (a ACLsByPrincipal) MarshalResourceAcls() ([]*sarama.ResourceAcls, error) {
	// convert to ACLsByResource for more efficiency
	abr := make(ACLsByResource)
	for p := range a {
		for _, acl := range a[p] {
			abr[acl.Resource] = append(abr[acl.Resource], acl)
		}
	}

	return abr.MarshalResourceAcls()
}

// MarshalResourceAcls converts ACLsByPrincipalAndResource to a list of sarama.ResourceAcls
func (a ACLsByPrincipalAndResource) MarshalResourceAcls() ([]*sarama.ResourceAcls, error) {
	// convert to ACLsByResource for more efficiency
	abr := make(ACLsByResource)
	for p := range a {
		for r := range a[p] {
			abr[r] = append(abr[r], a[p][r]...)
		}
	}

	return abr.MarshalResourceAcls()
}

// MarshalResourceAcls converts ACLsByResource to a list of sarama.ResourceAcls
func (a ACLsByResource) MarshalResourceAcls() ([]*sarama.ResourceAcls, error) {
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

// MarshalResourceAcls converts ACLsByResourceAndPrincipal to a list of sarama.ResourceAcls
func (a ACLsByResourceAndPrincipal) MarshalResourceAcls() ([]*sarama.ResourceAcls, error) {
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
			Resource: Resource{
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

// UnmarshalResourceAcls converts a list of sarama.ResourceAcls to ACLsByResource
func (a ACLsByResource) UnmarshalResourceAcls(rACLs []sarama.ResourceAcls) error {
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

// UnmarshalResourceAcls converts a list of sarama.ResourceAcls to ACLsByResourceAndPrincipal
func (a ACLsByResourceAndPrincipal) UnmarshalResourceAcls(rACLs []*sarama.ResourceAcls) error {
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

func (a ACLsByPrincipal) UnmarshalResourceAcls(rACLs []*sarama.ResourceAcls) error {
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

func (a ACLsByPrincipalAndResource) UnmarshalResourceAcls(rACLs []*sarama.ResourceAcls) error {
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
