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
	"github.com/crowdstrike/gofalcon/falcon/client/it_automation"
)

var (
	// ITAutomationAllowEmptyValues ...
	ITAutomationAllowEmptyValues = []string{"tags."}
)

// ITAutomationTaskGenerator ...
type ITAutomationTaskGenerator struct {
	CrowdStrikeService
}

func (g *ITAutomationTaskGenerator) createResource(taskID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		taskID,
		fmt.Sprintf("it_automation_task_%s", strings.ReplaceAll(taskID, "-", "_")),
		"crowdstrike_it_automation_task",
		"crowdstrike",
		ITAutomationAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from CrowdStrike API,
// from each IT automation task create 1 TerraformResource.
func (g *ITAutomationTaskGenerator) InitResources() error {
	client := g.Args["client"].(*client.CrowdStrikeAPISpecification)

	// Check if specific tasks are requested via filter
	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("it_automation_tasks") {
			for _, value := range filter.AcceptableValues {
				resources = append(resources, g.createResource(value))
			}
		}
	}

	if len(resources) > 0 {
		g.Resources = resources
		return nil
	}

	// Query all IT automation tasks using the correct API
	queryParams := &it_automation.ITAutomationSearchTasksParams{
		Context: context.Background(),
	}

	resp, err := client.ItAutomation.ITAutomationSearchTasks(queryParams)
	if err != nil {
		return fmt.Errorf("failed to query IT automation tasks: %v", err)
	}

	if resp.Payload == nil || resp.Payload.Resources == nil {
		// No tasks found - this is valid, return empty list
		g.Resources = []terraformutils.Resource{}
		return nil
	}

	// Convert API response to resource IDs - Resources are already strings (IDs)
	resources = g.createResources(resp.Payload.Resources)
	g.Resources = resources
	return nil
}

func (g *ITAutomationTaskGenerator) createResources(taskIDs []string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, taskID := range taskIDs {
		if taskID != "" {
			resources = append(resources, g.createResource(taskID))
		}
	}
	return resources
}

// ITAutomationTaskGroupGenerator ...
type ITAutomationTaskGroupGenerator struct {
	CrowdStrikeService
}

func (g *ITAutomationTaskGroupGenerator) createResource(taskGroupID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		taskGroupID,
		fmt.Sprintf("it_automation_task_group_%s", strings.ReplaceAll(taskGroupID, "-", "_")),
		"crowdstrike_it_automation_task_group",
		"crowdstrike",
		ITAutomationAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from CrowdStrike API,
// from each IT automation task group create 1 TerraformResource.
func (g *ITAutomationTaskGroupGenerator) InitResources() error {
	client := g.Args["client"].(*client.CrowdStrikeAPISpecification)

	// Check if specific task groups are requested via filter
	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("it_automation_task_groups") {
			for _, value := range filter.AcceptableValues {
				resources = append(resources, g.createResource(value))
			}
		}
	}

	if len(resources) > 0 {
		g.Resources = resources
		return nil
	}

	// Query all IT automation task groups using the correct API
	queryParams := &it_automation.ITAutomationSearchTaskGroupsParams{
		Context: context.Background(),
	}

	resp, err := client.ItAutomation.ITAutomationSearchTaskGroups(queryParams)
	if err != nil {
		return fmt.Errorf("failed to query IT automation task groups: %v", err)
	}

	if resp.Payload == nil || resp.Payload.Resources == nil {
		// No task groups found - this is valid, return empty list
		g.Resources = []terraformutils.Resource{}
		return nil
	}

	// Convert API response to resource IDs - Resources are already strings (IDs)
	resources = g.createResources(resp.Payload.Resources)
	g.Resources = resources
	return nil
}

func (g *ITAutomationTaskGroupGenerator) createResources(taskGroupIDs []string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, taskGroupID := range taskGroupIDs {
		if taskGroupID != "" {
			resources = append(resources, g.createResource(taskGroupID))
		}
	}
	return resources
}

// ITAutomationPolicyGenerator ...
type ITAutomationPolicyGenerator struct {
	CrowdStrikeService
}

func (g *ITAutomationPolicyGenerator) createResource(policyID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		policyID,
		fmt.Sprintf("it_automation_policy_%s", strings.ReplaceAll(policyID, "-", "_")),
		"crowdstrike_it_automation_policy",
		"crowdstrike",
		ITAutomationAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from CrowdStrike API,
// from each IT automation policy create 1 TerraformResource.
func (g *ITAutomationPolicyGenerator) InitResources() error {
	client := g.Args["client"].(*client.CrowdStrikeAPISpecification)

	// Check if specific policies are requested via filter
	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("it_automation_policies") {
			for _, value := range filter.AcceptableValues {
				resources = append(resources, g.createResource(value))
			}
		}
	}

	if len(resources) > 0 {
		g.Resources = resources
		return nil
	}

	// Query all IT automation policies using the correct API
	queryParams := &it_automation.ITAutomationQueryPoliciesParams{
		Context: context.Background(),
	}

	resp, err := client.ItAutomation.ITAutomationQueryPolicies(queryParams)
	if err != nil {
		return fmt.Errorf("failed to query IT automation policies: %v", err)
	}

	if resp.Payload == nil || resp.Payload.Resources == nil {
		// No policies found - this is valid, return empty list
		g.Resources = []terraformutils.Resource{}
		return nil
	}

	// Get detailed information for each policy
	if len(resp.Payload.Resources) > 0 {
		getParams := &it_automation.ITAutomationGetPoliciesParams{
			Context: context.Background(),
			Ids:     resp.Payload.Resources,
		}

		detailResp, err := client.ItAutomation.ITAutomationGetPolicies(getParams)
		if err != nil {
			return fmt.Errorf("failed to get IT automation policy details: %v", err)
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

func (g *ITAutomationPolicyGenerator) createResources(policyIDs []string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, policyID := range policyIDs {
		if policyID != "" {
			resources = append(resources, g.createResource(policyID))
		}
	}
	return resources
}
