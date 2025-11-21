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
	"github.com/crowdstrike/gofalcon/falcon/client/host_group"
)

var (
	// HostGroupAllowEmptyValues ...
	HostGroupAllowEmptyValues = []string{"tags."}
)

// HostGroupGenerator ...
type HostGroupGenerator struct {
	CrowdStrikeService
}

func (g *HostGroupGenerator) createResources(hostGroupIDs []string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, hostGroupID := range hostGroupIDs {
		if hostGroupID != "" {
			resources = append(resources, g.createResource(hostGroupID))
		}
	}
	return resources
}

func (g *HostGroupGenerator) createResource(hostGroupID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		hostGroupID,
		fmt.Sprintf("host_group_%s", strings.ReplaceAll(hostGroupID, "-", "_")),
		"crowdstrike_host_group",
		"crowdstrike",
		HostGroupAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from CrowdStrike API,
// from each host group create 1 TerraformResource.
// Need Host Group ID as ID for terraform resource
func (g *HostGroupGenerator) InitResources() error {
	client := g.Args["client"].(*client.CrowdStrikeAPISpecification)

	// Check if specific host groups are requested via filter
	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("host_groups") {
			for _, value := range filter.AcceptableValues {
				resources = append(resources, g.createResource(value))
			}
		}
	}

	if len(resources) > 0 {
		g.Resources = resources
		return nil
	}

	// Query all host groups using the correct API
	queryParams := &host_group.QueryHostGroupsParams{
		Context: context.Background(),
	}

	resp, err := client.HostGroup.QueryHostGroups(queryParams)
	if err != nil {
		return fmt.Errorf("failed to query host groups: %v", err)
	}

	if resp.Payload == nil || resp.Payload.Resources == nil {
		// No host groups found - this is valid, return empty list
		g.Resources = []terraformutils.Resource{}
		return nil
	}

	// Get detailed information for each host group
	if len(resp.Payload.Resources) > 0 {
		getParams := &host_group.GetHostGroupsParams{
			Context: context.Background(),
			Ids:     resp.Payload.Resources,
		}

		detailResp, err := client.HostGroup.GetHostGroups(getParams)
		if err != nil {
			return fmt.Errorf("failed to get host group details: %v", err)
		}

		if detailResp.Payload != nil && detailResp.Payload.Resources != nil {
			// Convert API response to resource IDs
			hostGroupIDs := make([]string, 0, len(detailResp.Payload.Resources))
			for _, hostGroup := range detailResp.Payload.Resources {
				if hostGroup.ID != nil {
					hostGroupIDs = append(hostGroupIDs, *hostGroup.ID)
				}
			}
			g.Resources = g.createResources(hostGroupIDs)
		}
	}

	return nil
}
