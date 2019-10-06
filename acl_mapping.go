package kafkalib

import "github.com/Shopify/sarama"

const (
	ResourceUnknown         = "Unknown"
	ResourceAny             = "Any"
	ResourceTopic           = "Topic"
	ResourceGroup           = "Group"
	ResourceCluster         = "Cluster"
	ResourceTransactionalID = "TransactionalID"

	OperationUnknown         = "Unknown"
	OperationAny             = "Any"
	OperationAll             = "All"
	OperationRead            = "Read"
	OperationWrite           = "Write"
	OperationCreate          = "Create"
	OperationDelete          = "Delete"
	OperationAlter           = "Alter"
	OperationDescribe        = "Describe"
	OperationClusterAction   = "ClusterAction"
	OperationDescribeConfigs = "DescribeConfigs"
	OperationAlterConfigs    = "AlterConfigs"
	OperationIdempotentWrite = "IdempotentWrite"

	PermissionUnknown = "unknown"
	PermissionAny     = "any"
	PermissionDeny    = "denied"
	PermissionAllow   = "allowed"

	PatternUnknown  = "Unknown"
	PatternAny      = "Any"
	PatternMatch    = "Match"
	PatternLiteral  = "Literal"
	PatternPrefixed = "Prefixed"
)

// maps for string conversion
var resourceTypeToString = map[sarama.AclResourceType]string{
	sarama.AclResourceUnknown:         ResourceUnknown,
	sarama.AclResourceAny:             ResourceAny,
	sarama.AclResourceTopic:           ResourceTopic,
	sarama.AclResourceGroup:           ResourceGroup,
	sarama.AclResourceCluster:         ResourceCluster,
	sarama.AclResourceTransactionalID: ResourceTransactionalID,
}

var operationToString = map[sarama.AclOperation]string{
	sarama.AclOperationUnknown:         OperationUnknown,
	sarama.AclOperationAny:             OperationAny,
	sarama.AclOperationAll:             OperationAll,
	sarama.AclOperationRead:            OperationRead,
	sarama.AclOperationWrite:           OperationWrite,
	sarama.AclOperationCreate:          OperationCreate,
	sarama.AclOperationDelete:          OperationDelete,
	sarama.AclOperationAlter:           OperationAlter,
	sarama.AclOperationDescribe:        OperationDescribe,
	sarama.AclOperationClusterAction:   OperationClusterAction,
	sarama.AclOperationDescribeConfigs: OperationDescribeConfigs,
	sarama.AclOperationAlterConfigs:    OperationAlterConfigs,
	sarama.AclOperationIdempotentWrite: OperationIdempotentWrite,
}

var permissionToString = map[sarama.AclPermissionType]string{
	sarama.AclPermissionUnknown: PermissionUnknown,
	sarama.AclPermissionAny:     PermissionAny,
	sarama.AclPermissionDeny:    PermissionDeny,
	sarama.AclPermissionAllow:   PermissionAllow,
}

var resourcePatternToString = map[sarama.AclResourcePatternType]string{
	sarama.AclPatternUnknown:  PatternUnknown,
	sarama.AclPatternAny:      PatternAny,
	sarama.AclPatternMatch:    PatternMatch,
	sarama.AclPatternLiteral:  PatternLiteral,
	sarama.AclPatternPrefixed: PatternPrefixed,
}

var resourceTypeToID = map[string]sarama.AclResourceType{
	ResourceUnknown:         sarama.AclResourceUnknown,
	ResourceAny:             sarama.AclResourceAny,
	ResourceTopic:           sarama.AclResourceTopic,
	ResourceGroup:           sarama.AclResourceGroup,
	ResourceCluster:         sarama.AclResourceCluster,
	ResourceTransactionalID: sarama.AclResourceTransactionalID,
}

var permissionToID = map[string]sarama.AclPermissionType{
	PermissionUnknown: sarama.AclPermissionUnknown,
	PermissionAny:     sarama.AclPermissionAny,
	PermissionDeny:    sarama.AclPermissionDeny,
	PermissionAllow:   sarama.AclPermissionAllow,
}

// OperationToID maps a string to a acl operation
var OperationToID = map[string]sarama.AclOperation{
	OperationUnknown:         sarama.AclOperationUnknown,
	OperationAny:             sarama.AclOperationAny,
	OperationAll:             sarama.AclOperationAll,
	OperationRead:            sarama.AclOperationRead,
	OperationWrite:           sarama.AclOperationWrite,
	OperationCreate:          sarama.AclOperationCreate,
	OperationDelete:          sarama.AclOperationDelete,
	OperationAlter:           sarama.AclOperationAlter,
	OperationDescribe:        sarama.AclOperationDescribe,
	OperationClusterAction:   sarama.AclOperationClusterAction,
	OperationDescribeConfigs: sarama.AclOperationDescribeConfigs,
	OperationAlterConfigs:    sarama.AclOperationAlterConfigs,
	OperationIdempotentWrite: sarama.AclOperationIdempotentWrite,
}

var resourcePatternToID = map[string]sarama.AclResourcePatternType{
	PatternUnknown:  sarama.AclPatternUnknown,
	PatternAny:      sarama.AclPatternAny,
	PatternMatch:    sarama.AclPatternMatch,
	PatternLiteral:  sarama.AclPatternLiteral,
	PatternPrefixed: sarama.AclPatternPrefixed,
}
