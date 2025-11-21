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
	"github.com/crowdstrike/gofalcon/falcon/client/prevention_policies"
)

var (
	// PreventionPolicyAllowEmptyValues ...
	PreventionPolicyAllowEmptyValues = []string{"tags."}
)

// PreventionPolicyGenerator ...
type PreventionPolicyGenerator struct {
	CrowdStrikeService
}

func (g *PreventionPolicyGenerator) createResources(policyIDs []string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, policyID := range policyIDs {
		if policyID != "" {
			resources = append(resources, g.createResource(policyID, "crowdstrike_prevention_policy"))
		}
	}
	return resources
}

func (g *PreventionPolicyGenerator) createResource(policyID, resourceType string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		policyID,
		fmt.Sprintf("prevention_policy_%s", strings.ReplaceAll(policyID, "-", "_")),
		resourceType,
		"crowdstrike",
		PreventionPolicyAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from CrowdStrike API,
// from each prevention policy create 1 TerraformResource.
// Need Prevention Policy ID as ID for terraform resource
func (g *PreventionPolicyGenerator) InitResources() error {
	client := g.Args["client"].(*client.CrowdStrikeAPISpecification)

	// Check if specific policies are requested via filter
	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("prevention_policies") {
			for _, value := range filter.AcceptableValues {
				resources = append(resources, g.createResource(value, "crowdstrike_prevention_policy"))
			}
		}
	}

	if len(resources) > 0 {
		g.Resources = resources
		return nil
	}

	// Query all prevention policies using the correct API
	queryParams := &prevention_policies.QueryCombinedPreventionPoliciesParams{
		Context: context.Background(),
	}

	resp, err := client.PreventionPolicies.QueryCombinedPreventionPolicies(queryParams)
	if err != nil {
		return fmt.Errorf("failed to query prevention policies: %v", err)
	}

	if resp.Payload == nil || resp.Payload.Resources == nil {
		// No prevention policies found - this is valid, return empty list
		g.Resources = []terraformutils.Resource{}
		return nil
	}

	// Convert API response to resources with proper resource types
	for _, policy := range resp.Payload.Resources {
		if policy.ID != nil {
			resourceType := "crowdstrike_prevention_policy"

			// Determine specific resource type based on platform
			if policy.PlatformName != nil {
				switch strings.ToLower(*policy.PlatformName) {
				case "windows":
					resourceType = "crowdstrike_prevention_policy_windows"
				case "mac":
					resourceType = "crowdstrike_prevention_policy_mac"
				case "linux":
					resourceType = "crowdstrike_prevention_policy_linux"
				}
			}

			resources = append(resources, g.createResource(*policy.ID, resourceType))
		}
	}

	g.Resources = resources
	return nil
}
