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
	// SensorVisibilityExclusionAllowEmptyValues ...
	SensorVisibilityExclusionAllowEmptyValues = []string{"tags."}
)

// SensorVisibilityExclusionGenerator ...
type SensorVisibilityExclusionGenerator struct {
	CrowdStrikeService
}

func (g *SensorVisibilityExclusionGenerator) createResources(exclusionIDs []string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, exclusionID := range exclusionIDs {
		if exclusionID != "" {
			resources = append(resources, g.createResource(exclusionID))
		}
	}
	return resources
}

func (g *SensorVisibilityExclusionGenerator) createResource(exclusionID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		exclusionID,
		fmt.Sprintf("sensor_visibility_exclusion_%s", strings.ReplaceAll(exclusionID, "-", "_")),
		"crowdstrike_sensor_visibility_exclusion",
		"crowdstrike",
		SensorVisibilityExclusionAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from CrowdStrike API,
// from each sensor visibility exclusion create 1 TerraformResource.
// Need Sensor Visibility Exclusion ID as ID for terraform resource
func (g *SensorVisibilityExclusionGenerator) InitResources() error {
	// Check if specific exclusions are requested via filter
	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("sensor_visibility_exclusions") {
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
	// TODO: Implement proper sensor visibility exclusion querying once we understand the gofalcon API structure
	g.Resources = []terraformutils.Resource{}
	return nil
}
