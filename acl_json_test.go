package kafkalib

import (
	"reflect"
	"testing"
)

func Test_acls_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		a       _acls
		want    []byte
		wantErr bool
	}{
		{
			name: "two acls",
			a: _acls{
				{
					Principal:      "User:test",
					PermissionType: PermissionAllow,
					Operation:      OperationAlter,
					Host:           "*",
					Resource: Resource{
						ResourceName: "test",
						ResourceType: ResourceTopic,
					},
				},
				{
					Principal:      "User:two",
					PermissionType: PermissionAllow,
					Operation:      OperationRead,
					Host:           "*",
					Resource: Resource{
						ResourceName: "test",
						ResourceType: ResourceTopic,
					},
				},
			},
			want:    []byte(string(`[{"principal":"User:test","permission_type":"Allow","operation":"Alter","host":"*","resource_name":"test","resource_type":"Topic"},{"principal":"User:two","permission_type":"Allow","operation":"Read","host":"*","resource_name":"test","resource_type":"Topic"}]`)),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("ACLs.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ACLs.MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}

func Test_abp_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		a       _abp
		want    []byte
		wantErr bool
	}{
		{
			name: "two acls",
			a: _abp{
				"User:test": []ACL{
					{
						Principal:      "User:test",
						PermissionType: PermissionAllow,
						Operation:      OperationAlter,
						Host:           "*",
						Resource: Resource{
							ResourceName: "test",
							ResourceType: ResourceTopic,
						},
					},
				},
				"User:two": []ACL{
					{
						Principal:      "User:two",
						PermissionType: PermissionAllow,
						Operation:      OperationRead,
						Host:           "*",
						Resource: Resource{
							ResourceName: "test",
							ResourceType: ResourceTopic,
						},
					},
				},
			},
			want:    []byte(`{"User:test":[{"principal":"User:test","permission_type":"Allow","operation":"Alter","host":"*","resource_name":"test","resource_type":"Topic"}],"User:two":[{"principal":"User:two","permission_type":"Allow","operation":"Read","host":"*","resource_name":"test","resource_type":"Topic"}]}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("_abp.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("_abp.MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestACLsByResource_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		a       ACLsByResource
		want    []byte
		wantErr bool
	}{
		{
			name: "two topics",
			a: ACLsByResource{
				Resource{
					ResourceName: "test",
					ResourceType: ResourceTopic,
				}: ACLs{
					{
						Principal:      "User:test",
						PermissionType: PermissionAllow,
						Operation:      OperationAlter,
						Host:           "*",
						Resource: Resource{
							ResourceName: "test",
							ResourceType: ResourceTopic,
						},
					},
					{
						Principal:      "User:two",
						PermissionType: PermissionAllow,
						Operation:      OperationRead,
						Host:           "*",
						Resource: Resource{
							ResourceName: "test",
							ResourceType: ResourceTopic,
						},
					},
				},
				Resource{
					ResourceName: "test2",
					ResourceType: ResourceTopic,
				}: ACLs{
					{
						Principal:      "User:test",
						PermissionType: PermissionAllow,
						Operation:      OperationAlter,
						Host:           "*",
						Resource: Resource{
							ResourceName: "test2",
							ResourceType: ResourceTopic,
						},
					},
				},
			},
			want:    []byte(`{"Topic_test":[{"principal":"User:test","permission_type":"Allow","operation":"Alter","host":"*","resource_name":"test","resource_type":"Topic"},{"principal":"User:two","permission_type":"Allow","operation":"Read","host":"*","resource_name":"test","resource_type":"Topic"}],"Topic_test2":[{"principal":"User:test","permission_type":"Allow","operation":"Alter","host":"*","resource_name":"test2","resource_type":"Topic"}]}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("_abr.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("_abr.MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}
