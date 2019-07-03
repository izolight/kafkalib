package entities_test

import (
	"encoding/json"
	. "github.com/izolight/kafkalib/entities"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/Shopify/sarama"
	"github.com/ghodss/yaml"
)

func TestACLMarhsalling(t *testing.T) {
	testAcl := ACLsByResourceType{
		"topic": ACLsByResource{
			"testtopic": []ACLByResource{
				{
					Principal: "User:testuser",
					ACLs: []ACL{
						{"*", "read", "allow"},
						{"*", "write", "allow"},
					},
				},
			},
		},
	}

	testAclPrincipal := ACLsByPrincipal{
		"User:testuser": []ACLByPrincipal{
			{
				Resource:     "testtopic",
				ResourceType: "topic",
				ACLs: []ACL{
					{"*", "read", "allow"},
					{"*", "write", "allow"},
				},
			},
		},
	}

	testAclJSON, err := json.MarshalIndent(testAcl, "", "  ")
	if err != nil {
		t.Error(err)
	}
	err = ioutil.WriteFile("./testdata/test.json", testAclJSON, 0644)
	if err != nil {
		t.Error(err)
	}
	testAclYAML, err := yaml.JSONToYAML(testAclJSON)
	if err != nil {
		t.Error(err)
	}
	err = ioutil.WriteFile("./testdata/test.yml", testAclYAML, 0644)
	if err != nil {
		t.Error(err)
	}

	testAclPrincipalJSON, err := json.MarshalIndent(testAclPrincipal, "", "  ")
	if err != nil {
		t.Error(err)
	}
	err = ioutil.WriteFile("./testdata/test-by-principal.json", testAclPrincipalJSON, 0644)
	if err != nil {
		t.Error(err)
	}
	testAclPrincipalYAML, err := yaml.JSONToYAML(testAclPrincipalJSON)
	if err != nil {
		t.Error(err)
	}
	err = ioutil.WriteFile("./testdata/test-by-principal.yml", testAclPrincipalYAML, 0644)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(testAclJSON, &testAcl)
	if err != nil {
		t.Error(err)
	}
	err = yaml.Unmarshal(testAclYAML, &testAcl)
	if err != nil {
		t.Error(err)
	}
	err = yaml.Unmarshal(testAclPrincipalYAML, &testAclPrincipal)
	if err != nil {
		t.Error(err)
	}
}

func TestACLsByPrincipal_MarshalSarama(t *testing.T) {
	tests := []struct {
		name    string
		a       ACLsByPrincipal
		want    []sarama.ResourceAcls
		wantErr bool
	}{
		{
			"simpleTest",
			ACLsByPrincipal{
				"User:Testuser": []ACLByPrincipal{
					{
						Resource:     "testtopic",
						ResourceType: "Topic",
						ACLs: []ACL{
							{
								Host:           "*",
								Operation:      "Read",
								PermissionType: "Allow",
							},
						},
					},
				},
			},
			[]sarama.ResourceAcls{
				{
					Resource: sarama.Resource{
						ResourceType: sarama.AclResourceTopic,
						ResourceName: "testtopic",
					},
					Acls: []*sarama.Acl{
						{
							Host:           "*",
							Operation:      sarama.AclOperationRead,
							PermissionType: sarama.AclPermissionAllow,
							Principal:      "User:Testuser",
						},
					},
				},
			},
			false,
		},
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

func TestACLsByPrincipal_UnmarshalSarama(t *testing.T) {
	type args struct {
		rAcls []sarama.ResourceAcls
	}
	tests := []struct {
		name    string
		want    ACLsByPrincipal
		args    args
		wantErr bool
	}{
		{
			"simpleTest",
			ACLsByPrincipal{
				"User:Testuser": []ACLByPrincipal{
					{
						Resource:     "testtopic",
						ResourceType: "Topic",
						ACLs: []ACL{
							{
								Host:           "*",
								Operation:      "Read",
								PermissionType: "Allow",
							},
							{
								Host:           "*",
								Operation:      "Write",
								PermissionType: "Allow",
							},
						},
					},
				},
			},
			args{
				[]sarama.ResourceAcls{
					{
						Resource: sarama.Resource{
							ResourceType: sarama.AclResourceTopic,
							ResourceName: "testtopic",
						},
						Acls: []*sarama.Acl{
							{
								Host:           "*",
								Operation:      sarama.AclOperationRead,
								PermissionType: sarama.AclPermissionAllow,
								Principal:      "User:Testuser",
							},
							{
								Host:           "*",
								Operation:      sarama.AclOperationWrite,
								PermissionType: sarama.AclPermissionAllow,
								Principal:      "User:Testuser",
							},
						},
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalSarama(tt.args.rAcls)
			if err != nil && !tt.wantErr {
				t.Errorf("ACLsByPrincipal.UnmarshalSarama() error = %v, wantErr %v", err, tt.wantErr)
			}
			for k := range tt.want {
				_, ok := got[k]
				if !ok {
					t.Errorf("ACLsByPrincipal.UnmarshalSarama() key %s not found in %s", k, got)
				}
				for i, a := range got[k] {
					if a.Resource != tt.want[k][i].Resource && a.ResourceType != tt.want[k][i].ResourceType {
						t.Errorf("ACLsByPrincipal.UnmarshalSarama() got different resource %s-%s, expected %s-%s",
							a.Resource, a.ResourceType, tt.want[k][i].Resource, tt.want[k][i].ResourceType)
					}
					for j, acl := range a.ACLs {
						if acl.Host != tt.want[k][i].ACLs[j].Host {
							t.Errorf("ACLsByPrincipal.UnmarshalSarama() got different Host %s, expected %s", acl.Host, tt.want[k][i].ACLs[j].Host)
						}
						if acl.PermissionType != tt.want[k][i].ACLs[j].PermissionType {
							t.Errorf("ACLsByPrincipal.UnmarshalSarama() got different Permission %s, expected %s", acl.PermissionType, tt.want[k][i].ACLs[j].PermissionType)
						}
						if acl.Operation != tt.want[k][i].ACLs[j].Operation {
							t.Errorf("ACLsByPrincipal.UnmarshalSarama() got different Operation %s, expected %s", acl.Operation, tt.want[k][i].ACLs[j].Operation)
						}
					}
				}
			}
		})
	}
}
