// Copyright 2023 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package guide

import (
	"sort"
	"strings"

	"github.com/cloudhut/connect-client"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/redpanda-data/console/backend/pkg/connector/model"
)

// DefaultGuide is the guide that is used if we haven't defined a guide for
type DefaultGuide struct {
	options Options
}

func NewDefaultGuide(opts ...Option) Guide {
	var o Options
	for _, opt := range opts {
		opt(&o)
	}
	return &DefaultGuide{options: o}
}

// ClassName implements Guide.ClassName().
func (g *DefaultGuide) ClassName() string {
	// The class name is used for guide matching, but the default guide is always the fallback guide
	// that is used when no other guides were matched. Hence no class name to be returned.
	return ""
}

func (g *DefaultGuide) ConsoleToKafkaConnect(configs map[string]any) map[string]any {
	for _, injectedVal := range g.options.injectedValues {
		if injectedVal.IsAuthoritative {
			// We are allowed to override existing user configs
			configs[injectedVal.Key] = injectedVal.Value
			continue
		}

		// We are not allowed to override existing user configs
		if _, exists := configs[injectedVal.Key]; !exists {
			configs[injectedVal.Key] = injectedVal.Value
		}
	}

	if g.options.consoleToKafkaConnectHookFn == nil {
		return configs
	} else {
		return g.options.consoleToKafkaConnectHookFn(configs)
	}
}

// KafkaConnectToConsole receives the response from the kafka connect cluster and returns
// the enriched validation response that is understood by Console. The Console validation response
// contains additional metadata.
func (g *DefaultGuide) KafkaConnectToConsole(response connect.ConnectorValidationResult) model.ValidationResponse {
	// 1. Extract all configs from the response and index them by their config key
	configs := make([]model.ConfigDefinition, len(response.Configs))
	configsByGroup := make(map[string][]model.ConfigDefinition)
	for i, config := range response.Configs {
		configDef := model.NewConfigDefinitionFromValidationResult(config)
		configs[i] = configDef

		group := ""
		if configDef.Definition.Group != nil {
			group = *configDef.Definition.Group
		}
		configsByGroup[group] = append(configsByGroup[group], configDef)
	}

	// 2. Sort grouped configs by their reported order
	for _, groupedDefs := range configsByGroup {
		slices.SortFunc(groupedDefs, func(a, b model.ConfigDefinition) bool {
			return a.Definition.Order < b.Definition.Order
		})
	}

	// 3. Order groups in order to achieve a stable order in the UI.
	// We try to order groups so that the most important configurations for the users show first.
	// The opinionated grouping is determined as follows:
	// - Common group first, then
	// - Groups that have the most important configurations show first, if multiple groups are equal:
	// - Order by group name asc
	importanceScoreByGroupName := make(map[string]int)
	for groupName, configDefs := range configsByGroup {
		importanceScoreByGroupName[groupName] = 0
		if strings.ToLower(groupName) == "common" {
			// Push the common group to the top
			importanceScoreByGroupName[groupName] += 100
		}
		for _, configDef := range configDefs {
			if configDef.Definition.Required || configDef.Definition.Importance == model.ConfigDefinitionImportanceHigh {
				importanceScoreByGroupName[groupName]++
			}
		}
	}
	groupNames := maps.Keys(importanceScoreByGroupName)
	// Sort by groupname asc
	slices.SortFunc(groupNames, func(a, b string) bool {
		return a < b
	})
	// Sort by number of required props
	sort.SliceStable(groupNames, func(a, b int) bool {
		grpNameA := groupNames[a]
		grpNameB := groupNames[b]
		return importanceScoreByGroupName[grpNameA] > importanceScoreByGroupName[grpNameB]
	})

	// 4. Convert all configs that are grouped by their group name into step groups
	stepGroups := make([]model.ValidationResponseStepGroup, 0, len(configsByGroup))
	for _, groupName := range groupNames {
		configDefs := configsByGroup[groupName]
		configKeys := make([]string, len(configDefs))
		for i, configDef := range configDefs {
			configKeys[i] = configDef.Definition.Name
		}

		sg := model.ValidationResponseStepGroup{
			Name:       groupName,
			ConfigKeys: configKeys,
		}
		stepGroups = append(stepGroups, sg)
	}

	validationResponse := model.ValidationResponse{
		Name:    response.Name,
		Configs: configs,
		Steps: []model.ValidationResponseStep{
			{
				Name:   "General",
				Groups: stepGroups,
			},
		},
	}

	if g.options.kafkaConnectToConsoleHookFn == nil {
		return validationResponse
	} else {
		return g.options.kafkaConnectToConsoleHookFn(response, validationResponse)
	}
}
