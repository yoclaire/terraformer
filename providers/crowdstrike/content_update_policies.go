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
	"github.com/crowdstrike/gofalcon/falcon/client/content_update_policies"
)

var (
	// ContentUpdatePolicyAllowEmptyValues ...
	ContentUpdatePolicyAllowEmptyValues = []string{"tags."}
)

// ContentUpdatePolicyGenerator ...
type ContentUpdatePolicyGenerator struct {
	CrowdStrikeService
}

func (g *ContentUpdatePolicyGenerator) createResources(policyIDs []string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, policyID := range policyIDs {
		if policyID != "" {
			resources = append(resources, g.createResource(policyID))
		}
	}
	return resources
}

func (g *ContentUpdatePolicyGenerator) createResource(policyID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		policyID,
		fmt.Sprintf("content_update_policy_%s", strings.ReplaceAll(policyID, "-", "_")),
		"crowdstrike_content_update_policy",
		"crowdstrike",
		ContentUpdatePolicyAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from CrowdStrike API,
// from each content update policy create 1 TerraformResource.
// Need Content Update Policy ID as ID for terraform resource
func (g *ContentUpdatePolicyGenerator) InitResources() error {
	client := g.Args["client"].(*client.CrowdStrikeAPISpecification)

	// Check if specific policies are requested via filter
	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("content_update_policies") {
			for _, value := range filter.AcceptableValues {
				resources = append(resources, g.createResource(value))
			}
		}
	}

	if len(resources) > 0 {
		g.Resources = resources
		return nil
	}

	// Query all content update policies using the correct API
	queryParams := &content_update_policies.QueryCombinedContentUpdatePoliciesParams{
		Context: context.Background(),
	}

	resp, err := client.ContentUpdatePolicies.QueryCombinedContentUpdatePolicies(queryParams)
	if err != nil {
		return fmt.Errorf("failed to query content update policies: %v", err)
	}

	if resp.Payload == nil || resp.Payload.Resources == nil {
		// No content update policies found - this is valid, return empty list
		g.Resources = []terraformutils.Resource{}
		return nil
	}

	// Convert API response to resource IDs
	policyIDs := make([]string, 0, len(resp.Payload.Resources))
	for _, policy := range resp.Payload.Resources {
		if policy.ID != nil {
			policyIDs = append(policyIDs, *policy.ID)
		}
	}

	g.Resources = g.createResources(policyIDs)
	return nil
}
