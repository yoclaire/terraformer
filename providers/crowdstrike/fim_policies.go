// Copyright 2018 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package crowdstrike

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/crowdstrike/gofalcon/falcon/client"
	"github.com/crowdstrike/gofalcon/falcon/client/filevantage"
)

var (
	// FIMPolicyAllowEmptyValues ...
	FIMPolicyAllowEmptyValues = []string{"tags."}
)

// FIMPolicyGenerator ...
type FIMPolicyGenerator struct {
	CrowdStrikeService
}

func (g *FIMPolicyGenerator) createResource(policyID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		policyID,
		fmt.Sprintf("fim_policy_%s", strings.ReplaceAll(policyID, "-", "_")),
		"crowdstrike_fim_policy",
		"crowdstrike",
		FIMPolicyAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from CrowdStrike API,
// from each FIM policy create 1 TerraformResource.
// Need FIM Policy ID as ID for terraform resource
func (g *FIMPolicyGenerator) InitResources() error {
	client := g.Args["client"].(*client.CrowdStrikeAPISpecification)

	// Check if specific policies are requested via filter
	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("fim_policies") {
			for _, value := range filter.AcceptableValues {
				resources = append(resources, g.createResource(value))
			}
		}
	}

	if len(resources) > 0 {
		g.Resources = resources
		return nil
	}

	// Query all filevantage policies using the correct API
	queryParams := &filevantage.QueryPoliciesParams{
		Context: context.Background(),
	}

	resp, err := client.Filevantage.QueryPolicies(queryParams)
	if err != nil {
		return fmt.Errorf("failed to query FIM policies: %v", err)
	}

	if resp.Payload == nil || resp.Payload.Resources == nil {
		// No FIM policies found - this is valid, return empty list
		g.Resources = []terraformutils.Resource{}
		return nil
	}

	// Get detailed information for each policy
	if len(resp.Payload.Resources) > 0 {
		getParams := &filevantage.GetPoliciesParams{
			Context: context.Background(),
			Ids:     resp.Payload.Resources,
		}

		detailResp, err := client.Filevantage.GetPolicies(getParams)
		if err != nil {
			return fmt.Errorf("failed to get FIM policy details: %v", err)
		}

		if detailResp.Payload != nil && detailResp.Payload.Resources != nil {
			// Convert API response to resource IDs
			policyIDs := make([]string, 0, len(detailResp.Payload.Resources))
			for _, policy := range detailResp.Payload.Resources {
				if policy.ID != nil {
					policyIDs = append(policyIDs, *policy.ID)
				}
			}
			resources = g.createResources(policyIDs)
		}
	}

	g.Resources = resources
	return nil
}

func (g *FIMPolicyGenerator) createResources(policyIDs []string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, policyID := range policyIDs {
		if policyID != "" {
			resources = append(resources, g.createResource(policyID))
		}
	}
	return resources
}

// FilevantageRuleGroupGenerator ...
type FilevantageRuleGroupGenerator struct {
	CrowdStrikeService
}

func (g *FilevantageRuleGroupGenerator) createResource(ruleGroupID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		ruleGroupID,
		fmt.Sprintf("filevantage_rule_group_%s", strings.ReplaceAll(ruleGroupID, "-", "_")),
		"crowdstrike_filevantage_rule_group",
		"crowdstrike",
		FIMPolicyAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from CrowdStrike API,
// from each Filevantage rule group create 1 TerraformResource.
// Need Filevantage Rule Group ID as ID for terraform resource
func (g *FilevantageRuleGroupGenerator) InitResources() error {
	client := g.Args["client"].(*client.CrowdStrikeAPISpecification)

	// Check if specific rule groups are requested via filter
	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("filevantage_rule_groups") {
			for _, value := range filter.AcceptableValues {
				resources = append(resources, g.createResource(value))
			}
		}
	}

	if len(resources) > 0 {
		g.Resources = resources
		return nil
	}

	// Query all filevantage rule groups using the correct API
	queryParams := &filevantage.QueryRuleGroupsParams{
		Context: context.Background(),
	}

	resp, err := client.Filevantage.QueryRuleGroups(queryParams)
	if err != nil {
		return fmt.Errorf("failed to query filevantage rule groups: %v", err)
	}

	if resp.Payload == nil || resp.Payload.Resources == nil {
		// No rule groups found - this is valid, return empty list
		g.Resources = []terraformutils.Resource{}
		return nil
	}

	// Get detailed information for each rule group
	if len(resp.Payload.Resources) > 0 {
		getParams := &filevantage.GetRuleGroupsParams{
			Context: context.Background(),
			Ids:     resp.Payload.Resources,
		}

		detailResp, err := client.Filevantage.GetRuleGroups(getParams)
		if err != nil {
			return fmt.Errorf("failed to get filevantage rule group details: %v", err)
		}

		if detailResp.Payload != nil && detailResp.Payload.Resources != nil {
			// Convert API response to resource IDs
			ruleGroupIDs := make([]string, 0, len(detailResp.Payload.Resources))
			for _, ruleGroup := range detailResp.Payload.Resources {
				if ruleGroup.ID != nil {
					ruleGroupIDs = append(ruleGroupIDs, *ruleGroup.ID)
				}
			}
			resources = g.createResources(ruleGroupIDs)
		}
	}

	g.Resources = resources
	return nil
}

func (g *FilevantageRuleGroupGenerator) createResources(ruleGroupIDs []string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, ruleGroupID := range ruleGroupIDs {
		if ruleGroupID != "" {
			resources = append(resources, g.createResource(ruleGroupID))
		}
	}
	return resources
}
