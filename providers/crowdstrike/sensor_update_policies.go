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
	// SensorUpdatePolicyAllowEmptyValues ...
	SensorUpdatePolicyAllowEmptyValues = []string{"tags."}
)

// SensorUpdatePolicyGenerator ...
type SensorUpdatePolicyGenerator struct {
	CrowdStrikeService
}

func (g *SensorUpdatePolicyGenerator) createResources(policyIDs []string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, policyID := range policyIDs {
		if policyID != "" {
			resources = append(resources, g.createResource(policyID))
		}
	}
	return resources
}

func (g *SensorUpdatePolicyGenerator) createResource(policyID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		policyID,
		fmt.Sprintf("sensor_update_policy_%s", strings.ReplaceAll(policyID, "-", "_")),
		"crowdstrike_sensor_update_policy",
		"crowdstrike",
		SensorUpdatePolicyAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from CrowdStrike API,
// from each sensor update policy create 1 TerraformResource.
// Need Sensor Update Policy ID as ID for terraform resource
func (g *SensorUpdatePolicyGenerator) InitResources() error {
	// Check if specific policies are requested via filter
	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("sensor_update_policies") {
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
	// TODO: Implement proper sensor update policy querying once we understand the gofalcon API structure
	g.Resources = []terraformutils.Resource{}
	return nil
}
