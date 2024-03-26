// Copyright 2022 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package console

import (
	"context"
	"fmt"
	"strings"

	"github.com/twmb/franz-go/pkg/kerr"
	"github.com/twmb/franz-go/pkg/kmsg"
)

// ACLPrincipalType is the type of principal.
// Example "User", or "RedpandaRole".
type ACLPrincipalType string

const (
	// ACLPrincipalTypeUnknown is the default unknown principal type.
	ACLPrincipalTypeUnknown ACLPrincipalType = "UNKNOWN"
	// ACLPrincipalTypeUser is the User principal type.
	ACLPrincipalTypeUser ACLPrincipalType = "USER"
	// ACLPrincipalTypeGroup is the Group principal type.
	ACLPrincipalTypeGroup ACLPrincipalType = "GROUP"
	// ACLPrincipalTypeRedpandaRole is the RedpandaRole principal type.
	ACLPrincipalTypeRedpandaRole ACLPrincipalType = "REDPANDA_ROLE"
)

// ACLOverview contains all acl resources along with the information whether an
// authorizer is enabled in the target cluster at all.
type ACLOverview struct {
	ACLResources        []*ACLResource `json:"aclResources"`
	IsAuthorizerEnabled bool           `json:"isAuthorizerEnabled"`

	// KafkaResponse is the original Kafka response. This is used by the connect
	// RPC API as this API has its own mapping logic.
	KafkaResponse *kmsg.DescribeACLsResponse `json:"-"`
}

// ACLResource is all information we get when listing ACLs
type ACLResource struct {
	ResourceType        string     `json:"resourceType"`
	ResourceName        string     `json:"resourceName"`
	ResourcePatternType string     `json:"resourcePatternType"`
	ACLs                []*ACLRule `json:"acls"`
}

// ACLRule describes a Kafka ACL rule with all it's properties.
type ACLRule struct {
	Principal      string           `json:"principal"`
	Host           string           `json:"host"`
	Operation      string           `json:"operation"`
	PermissionType string           `json:"permissionType"`
	PrincipalName  string           `json:"principalName"`
	PrincipalType  ACLPrincipalType `json:"principalType"`
}

// ListAllACLs returns a list of all stored ACLs.
func (s *Service) ListAllACLs(ctx context.Context, req kmsg.DescribeACLsRequest) (*ACLOverview, error) {
	aclResponses, err := s.kafkaSvc.ListACLs(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get ACLs from Kafka: %w", err)
	}

	kafkaErr := kerr.TypedErrorForCode(aclResponses.ErrorCode)
	if kafkaErr != nil {
		if kafkaErr == kerr.SecurityDisabled {
			return &ACLOverview{
				ACLResources:        nil,
				IsAuthorizerEnabled: false,
			}, nil
		}
		return nil, fmt.Errorf("failed to get ACLs from Kafka: %w", kafkaErr)
	}

	resources := make([]*ACLResource, len(aclResponses.Resources))
	for i, aclResponse := range aclResponses.Resources {
		overview := &ACLResource{
			ResourceType:        aclResponse.ResourceType.String(),
			ResourceName:        aclResponse.ResourceName,
			ResourcePatternType: aclResponse.ResourcePatternType.String(),
			ACLs:                nil,
		}

		acls := make([]*ACLRule, len(aclResponse.ACLs))
		for j, acl := range aclResponse.ACLs {
			acls[j] = &ACLRule{
				Principal:      acl.Principal,
				Host:           acl.Host,
				Operation:      acl.Operation.String(),
				PermissionType: acl.PermissionType.String(),
			}

			p := strings.ToLower(acl.Principal)

			switch {
			case strings.HasPrefix(p, "user:"):
				acls[j].PrincipalType = ACLPrincipalTypeUser
				acls[j].PrincipalName = p[strings.IndexRune(p, ':')+1:]
			case strings.HasPrefix(p, "group:"):
				acls[j].PrincipalType = ACLPrincipalTypeGroup
				acls[j].PrincipalName = p[strings.IndexRune(p, ':')+1:]
			case strings.HasPrefix(p, "redpandarole:"):
				acls[j].PrincipalType = ACLPrincipalTypeRedpandaRole
				acls[j].PrincipalName = p[strings.IndexRune(p, ':')+1:]
			default:
				acls[j].PrincipalType = ACLPrincipalTypeUnknown
				acls[j].PrincipalName = p
			}
		}
		overview.ACLs = acls
		resources[i] = overview
	}

	return &ACLOverview{
		ACLResources:        resources,
		IsAuthorizerEnabled: true,
		KafkaResponse:       aclResponses,
	}, nil
}
