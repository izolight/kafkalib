package kafkalib

import "github.com/Shopify/sarama"

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

var resourcePatternToString = map[sarama.AclResourcePatternType]string {
	sarama.AclPatternUnknown: "Unknown",
	sarama.AclPatternAny: "Any",
	sarama.AclPatternMatch: "Match",
	sarama.AclPatternLiteral: "Literal",
	sarama.AclPatternPrefixed: "Prefixed",
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

var resourcePatternToID = map[string]sarama.AclResourcePatternType{
	"Unknown": sarama.AclPatternUnknown,
	"Any": sarama.AclPatternAny,
	"Match": sarama.AclPatternMatch,
	"Literal": sarama.AclPatternLiteral,
	"Prefixed": sarama.AclPatternPrefixed,
}