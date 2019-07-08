package kafkalib

import (
	"fmt"
	"github.com/Shopify/sarama"
	"regexp"
)

func NewAclFilter(filter string) (*sarama.AclFilter, error) {
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

	af := &sarama.AclFilter{}
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

	return af, nil
}