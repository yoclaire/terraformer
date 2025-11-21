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
	// TODO: Implement FIM policy discovery using the appropriate gofalcon client
	// This is a placeholder implementation
	return fmt.Errorf("FIM policy resource generator not yet implemented")
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
	// TODO: Implement Filevantage rule group discovery using the appropriate gofalcon client
	// This is a placeholder implementation
	return fmt.Errorf("Filevantage rule group resource generator not yet implemented")
}
