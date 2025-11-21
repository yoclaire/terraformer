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
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
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

	// For now, return empty resources as we need to determine the correct API method
	// TODO: Implement proper host group querying once we understand the gofalcon API structure
	g.Resources = []terraformutils.Resource{}
	return nil
}
