package kafkalib

import (
	"encoding/json"
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
				t.Errorf("ACLsByPrincipalAndResource.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ACLsByPrincipalAndResource.MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestACLsByPrincipalAndResource_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		a       ACLsByPrincipalAndResource
		want    []byte
		wantErr bool
	}{
		{
			name: "three acls over two principals and 2 resources",
			a: ACLsByPrincipalAndResource{
				"User:test": map[Resource]ACLs{
					Resource{
						ResourceName: "test",
						ResourceType: ResourceTopic,
					}: {
						ACL{
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
					Resource{
						ResourceName: "test-2",
						ResourceType: ResourceTopic,
					}: {
						ACL{
							Principal:      "User:test",
							PermissionType: PermissionAllow,
							Operation:      OperationRead,
							Host:           "*",
							Resource: Resource{
								ResourceName: "test-2",
								ResourceType: ResourceTopic,
							},
						},
					},
				},
				"User:two": map[Resource]ACLs{
					Resource{
						ResourceName: "test",
						ResourceType: ResourceTopic,
					}: {
						ACL{
							Principal:      "User:two",
							PermissionType: PermissionAllow,
							Operation:      OperationDescribe,
							Host:           "*",
							Resource: Resource{
								ResourceName: "test",
								ResourceType: ResourceTopic,
							},
						},
					},
				},
			},
			want:    []byte(`{"User:test":[{"Topic_test":[{"principal":"User:test","permission_type":"Allow","operation":"Alter","host":"*","resource_name":"test","resource_type":"Topic"}]},{"Topic_test-2":[{"principal":"User:test","permission_type":"Allow","operation":"Read","host":"*","resource_name":"test-2","resource_type":"Topic"}]}],"User:two":[{"Topic_test":[{"principal":"User:two","permission_type":"Allow","operation":"Describe","host":"*","resource_name":"test","resource_type":"Topic"}]}]}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("ACLsByPrincipalAndResource.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !json.Valid(got) {
				t.Errorf("ACLsByPrincipalAndResource.MarshalJSON() is invalid = %s, want like this %s", got, tt.want)
			}
		})
	}
}

func TestACLsByResourceAndPrincipal_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		a       ACLsByResourceAndPrincipal
		want    []byte
		wantErr bool
	}{
		{
			name: "three acls over two principals and 2 resources",
			a: ACLsByResourceAndPrincipal{
				Resource{
					ResourceName: "test",
					ResourceType: ResourceTopic,
				}: map[string]ACLs{
					"User:test": {
						ACL{
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
					"User:two": {
						ACL{
							Principal:      "User:two",
							PermissionType: PermissionAllow,
							Operation:      OperationDescribe,
							Host:           "*",
							Resource: Resource{
								ResourceName: "test",
								ResourceType: ResourceTopic,
							},
						},
					},
				},
				Resource{
					ResourceName: "test-2",
					ResourceType: ResourceTopic,
				}: map[string]ACLs{
					"User:test": {
						ACL{
							Principal:      "User:test",
							PermissionType: PermissionAllow,
							Operation:      OperationRead,
							Host:           "*",
							Resource: Resource{
								ResourceName: "test-2",
								ResourceType: ResourceTopic,
							},
						},
					},
				},
			},
			want: []byte(`{"Topic_test":[{"User:test":[{"principal":"User:test","permission_type":"Allow","operation":"Alter","host":"*","resource_name":"test","resource_type":"Topic"}]},{"User:two":[{"principal":"User:two","permission_type":"Allow","operation":"Describe","host":"*","resource_name":"test","resource_type":"Topic"}]}],"Topic_test-2":[{"User:test":[{"principal":"User:test","permission_type":"Allow","operation":"Read","host":"*","resource_name":"test-2","resource_type":"Topic"}]}]}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("ACLsByResourceAndPrincipal.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !json.Valid(got) {
				t.Errorf("ACLsByResourceAndPrincipal.MarshalJSON() is invalid = %s, want like this %s", got, tt.want)
			}
		})
	}
}
