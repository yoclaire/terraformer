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
	// CloudResourceAllowEmptyValues ...
	CloudResourceAllowEmptyValues = []string{"tags."}
)

// CloudAWSAccountGenerator ...
type CloudAWSAccountGenerator struct {
	CrowdStrikeService
}

func (g *CloudAWSAccountGenerator) createResource(accountID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		accountID,
		fmt.Sprintf("cloud_aws_account_%s", strings.ReplaceAll(accountID, "-", "_")),
		"crowdstrike_cloud_aws_account",
		"crowdstrike",
		CloudResourceAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from CrowdStrike API,
// from each cloud AWS account create 1 TerraformResource.
func (g *CloudAWSAccountGenerator) InitResources() error {
	// TODO: Implement cloud AWS account discovery using the appropriate gofalcon client
	return fmt.Errorf("cloud AWS account resource generator not yet implemented")
}

// CloudAzureTenantGenerator ...
type CloudAzureTenantGenerator struct {
	CrowdStrikeService
}

func (g *CloudAzureTenantGenerator) createResource(tenantID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		tenantID,
		fmt.Sprintf("cloud_azure_tenant_%s", strings.ReplaceAll(tenantID, "-", "_")),
		"crowdstrike_cloud_azure_tenant",
		"crowdstrike",
		CloudResourceAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from CrowdStrike API,
// from each cloud Azure tenant create 1 TerraformResource.
func (g *CloudAzureTenantGenerator) InitResources() error {
	// TODO: Implement cloud Azure tenant discovery using the appropriate gofalcon client
	return fmt.Errorf("cloud Azure tenant resource generator not yet implemented")
}

// CloudSecurityCustomRuleGenerator ...
type CloudSecurityCustomRuleGenerator struct {
	CrowdStrikeService
}

func (g *CloudSecurityCustomRuleGenerator) createResource(ruleID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		ruleID,
		fmt.Sprintf("cloud_security_custom_rule_%s", strings.ReplaceAll(ruleID, "-", "_")),
		"crowdstrike_cloud_security_custom_rule",
		"crowdstrike",
		CloudResourceAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from CrowdStrike API,
// from each cloud security custom rule create 1 TerraformResource.
func (g *CloudSecurityCustomRuleGenerator) InitResources() error {
	// TODO: Implement cloud security custom rule discovery using the appropriate gofalcon client
	return fmt.Errorf("cloud security custom rule resource generator not yet implemented")
}

// CloudComplianceFrameworkGenerator ...
type CloudComplianceFrameworkGenerator struct {
	CrowdStrikeService
}

func (g *CloudComplianceFrameworkGenerator) createResource(frameworkID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		frameworkID,
		fmt.Sprintf("cloud_compliance_framework_%s", strings.ReplaceAll(frameworkID, "-", "_")),
		"crowdstrike_cloud_compliance_custom_framework",
		"crowdstrike",
		CloudResourceAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from CrowdStrike API,
// from each cloud compliance framework create 1 TerraformResource.
func (g *CloudComplianceFrameworkGenerator) InitResources() error {
	// TODO: Implement cloud compliance framework discovery using the appropriate gofalcon client
	return fmt.Errorf("cloud compliance framework resource generator not yet implemented")
}

// CloudGroupGenerator ...
type CloudGroupGenerator struct {
	CrowdStrikeService
}

func (g *CloudGroupGenerator) createResource(groupID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		groupID,
		fmt.Sprintf("cloud_group_%s", strings.ReplaceAll(groupID, "-", "_")),
		"crowdstrike_cloud_group",
		"crowdstrike",
		CloudResourceAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from CrowdStrike API,
// from each cloud group create 1 TerraformResource.
func (g *CloudGroupGenerator) InitResources() error {
	// TODO: Implement cloud group discovery using the appropriate gofalcon client
	return fmt.Errorf("cloud group resource generator not yet implemented")
}
