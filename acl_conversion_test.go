package kafkalib

import (
	"reflect"
	"testing"

	"github.com/Shopify/sarama"
)

func TestACL_MarshalSaramaRACL(t *testing.T) {
	type fields struct {
		Principal      string
		PermissionType string
		Operation      string
		Host           string
		Resource       Resource
	}
	tests := []struct {
		name    string
		fields  fields
		want    *sarama.ResourceAcls
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ACL{
				Principal:      tt.fields.Principal,
				PermissionType: tt.fields.PermissionType,
				Operation:      tt.fields.Operation,
				Host:           tt.fields.Host,
				Resource:       tt.fields.Resource,
			}
			got, err := a.MarshalSaramaRACL()
			if (err != nil) != tt.wantErr {
				t.Errorf("ACL.MarshalSaramaRACL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ACL.MarshalSaramaRACL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestACL_MarshalSaramaACL(t *testing.T) {
	type fields struct {
		Principal      string
		PermissionType string
		Operation      string
		Host           string
		Resource       Resource
	}
	tests := []struct {
		name    string
		fields  fields
		want    *sarama.Acl
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ACL{
				Principal:      tt.fields.Principal,
				PermissionType: tt.fields.PermissionType,
				Operation:      tt.fields.Operation,
				Host:           tt.fields.Host,
				Resource:       tt.fields.Resource,
			}
			got, err := a.MarshalSaramaACL()
			if (err != nil) != tt.wantErr {
				t.Errorf("ACL.MarshalSaramaACL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ACL.MarshalSaramaACL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestACLs_MarshalSaramaPerResource(t *testing.T) {
	tests := []struct {
		name    string
		a       ACLs
		want    *sarama.ResourceAcls
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.MarshalSaramaPerResource()
			if (err != nil) != tt.wantErr {
				t.Errorf("ACLs.MarshalSaramaPerResource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ACLs.MarshalSaramaPerResource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestACLsByPrincipal_MarshalSarama(t *testing.T) {
	tests := []struct {
		name    string
		a       ACLsByPrincipal
		want    []*sarama.ResourceAcls
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.MarshalSarama()
			if (err != nil) != tt.wantErr {
				t.Errorf("ACLsByPrincipal.MarshalSarama() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ACLsByPrincipal.MarshalSarama() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestACLsByPrincipalAndResource_MarshalSarama(t *testing.T) {
	tests := []struct {
		name    string
		a       ACLsByPrincipalAndResource
		want    []*sarama.ResourceAcls
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.MarshalSarama()
			if (err != nil) != tt.wantErr {
				t.Errorf("ACLsByPrincipalAndResource.MarshalSarama() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ACLsByPrincipalAndResource.MarshalSarama() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestACLsByResource_MarshalSarama(t *testing.T) {
	tests := []struct {
		name    string
		a       ACLsByResource
		want    []*sarama.ResourceAcls
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.MarshalSarama()
			if (err != nil) != tt.wantErr {
				t.Errorf("ACLsByResource.MarshalSarama() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ACLsByResource.MarshalSarama() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestACLsByResourceAndPrincipal_MarshalSarama(t *testing.T) {
	tests := []struct {
		name    string
		a       ACLsByResourceAndPrincipal
		want    []*sarama.ResourceAcls
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.MarshalSarama()
			if (err != nil) != tt.wantErr {
				t.Errorf("ACLsByResourceAndPrincipal.MarshalSarama() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ACLsByResourceAndPrincipal.MarshalSarama() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestACLs_UnmarshalSarama(t *testing.T) {
	type args struct {
		rACLs *sarama.ResourceAcls
	}
	tests := []struct {
		name    string
		a       *ACLs
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.UnmarshalSarama(tt.args.rACLs); (err != nil) != tt.wantErr {
				t.Errorf("ACLs.UnmarshalSarama() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestACL_UnmarshalSarama(t *testing.T) {
	type fields struct {
		Principal      string
		PermissionType string
		Operation      string
		Host           string
		Resource       Resource
	}
	type args struct {
		acl *sarama.Acl
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ACL{
				Principal:      tt.fields.Principal,
				PermissionType: tt.fields.PermissionType,
				Operation:      tt.fields.Operation,
				Host:           tt.fields.Host,
				Resource:       tt.fields.Resource,
			}
			if err := a.UnmarshalSarama(tt.args.acl); (err != nil) != tt.wantErr {
				t.Errorf("ACL.UnmarshalSarama() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestACLsByResource_UnmarshalSarama(t *testing.T) {
	type args struct {
		rACLs []*sarama.ResourceAcls
	}
	tests := []struct {
		name    string
		a       ACLsByResource
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.UnmarshalSarama(tt.args.rACLs); (err != nil) != tt.wantErr {
				t.Errorf("ACLsByResource.UnmarshalSarama() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestACLsByResourceAndPrincipal_UnmarshalSarama(t *testing.T) {
	type args struct {
		rACLs []*sarama.ResourceAcls
	}
	tests := []struct {
		name    string
		a       ACLsByResourceAndPrincipal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.UnmarshalSarama(tt.args.rACLs); (err != nil) != tt.wantErr {
				t.Errorf("ACLsByResourceAndPrincipal.UnmarshalSarama() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestACLsByPrincipal_UnmarshalSarama(t *testing.T) {
	type args struct {
		rACLs []*sarama.ResourceAcls
	}
	tests := []struct {
		name    string
		a       ACLsByPrincipal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.UnmarshalSarama(tt.args.rACLs); (err != nil) != tt.wantErr {
				t.Errorf("ACLsByPrincipal.UnmarshalSarama() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestACLsByPrincipalAndResource_UnmarshalSarama(t *testing.T) {
	type args struct {
		rACLs []*sarama.ResourceAcls
	}
	tests := []struct {
		name    string
		a       ACLsByPrincipalAndResource
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.UnmarshalSarama(tt.args.rACLs); (err != nil) != tt.wantErr {
				t.Errorf("ACLsByPrincipalAndResource.UnmarshalSarama() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
