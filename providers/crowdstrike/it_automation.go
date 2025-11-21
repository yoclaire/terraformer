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
	// TODO: Implement IT automation task discovery using the appropriate gofalcon client
	return fmt.Errorf("IT automation task resource generator not yet implemented")
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
	// TODO: Implement IT automation task group discovery using the appropriate gofalcon client
	return fmt.Errorf("IT automation task group resource generator not yet implemented")
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
	// TODO: Implement IT automation policy discovery using the appropriate gofalcon client
	return fmt.Errorf("IT automation policy resource generator not yet implemented")
}
