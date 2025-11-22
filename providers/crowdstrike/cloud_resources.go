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
	"github.com/crowdstrike/gofalcon/falcon/client/cloud_policies"
	"github.com/crowdstrike/gofalcon/falcon/client/cloud_security"
	"github.com/crowdstrike/gofalcon/falcon/client/cspm_registration"
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
	client := g.Args["client"].(*client.CrowdStrikeAPISpecification)

	// Check if specific accounts are requested via filter
	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("cloud_aws_accounts") {
			for _, value := range filter.AcceptableValues {
				resources = append(resources, g.createResource(value))
			}
		}
	}

	if len(resources) > 0 {
		g.Resources = resources
		return nil
	}

	// Query all cloud AWS accounts using the correct API
	queryParams := &cspm_registration.GetCSPMAwsAccountParams{
		Context: context.Background(),
	}

	resp, _, err := client.CspmRegistration.GetCSPMAwsAccount(queryParams)
	if err != nil {
		return fmt.Errorf("failed to query cloud AWS accounts: %v", err)
	}

	if resp.Payload == nil || resp.Payload.Resources == nil {
		// No AWS accounts found - this is valid, return empty list
		g.Resources = []terraformutils.Resource{}
		return nil
	}

	// Convert API response to resource IDs - Resources contain AccountID strings
	accountIDs := make([]string, 0, len(resp.Payload.Resources))
	for _, account := range resp.Payload.Resources {
		if account.AccountID != "" {
			accountIDs = append(accountIDs, account.AccountID)
		}
	}

	resources = g.createResources(accountIDs)
	g.Resources = resources
	return nil
}

func (g *CloudAWSAccountGenerator) createResources(accountIDs []string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, accountID := range accountIDs {
		if accountID != "" {
			resources = append(resources, g.createResource(accountID))
		}
	}
	return resources
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
	client := g.Args["client"].(*client.CrowdStrikeAPISpecification)

	// Check if specific tenants are requested via filter
	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("cloud_azure_tenants") {
			for _, value := range filter.AcceptableValues {
				resources = append(resources, g.createResource(value))
			}
		}
	}

	if len(resources) > 0 {
		g.Resources = resources
		return nil
	}

	// Query all cloud Azure tenants using the correct API
	queryParams := &cspm_registration.GetCSPMAzureAccountParams{
		Context: context.Background(),
	}

	resp, _, err := client.CspmRegistration.GetCSPMAzureAccount(queryParams)
	if err != nil {
		return fmt.Errorf("failed to query cloud Azure tenants: %v", err)
	}

	if resp.Payload == nil || resp.Payload.Resources == nil {
		// No Azure tenants found - this is valid, return empty list
		g.Resources = []terraformutils.Resource{}
		return nil
	}

	// Convert API response to resource IDs - Resources contain TenantID strings
	tenantIDs := make([]string, 0, len(resp.Payload.Resources))
	for _, tenant := range resp.Payload.Resources {
		if tenant.TenantID != "" {
			tenantIDs = append(tenantIDs, tenant.TenantID)
		}
	}

	resources = g.createResources(tenantIDs)
	g.Resources = resources
	return nil
}

func (g *CloudAzureTenantGenerator) createResources(tenantIDs []string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, tenantID := range tenantIDs {
		if tenantID != "" {
			resources = append(resources, g.createResource(tenantID))
		}
	}
	return resources
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
	client := g.Args["client"].(*client.CrowdStrikeAPISpecification)

	// Check if specific rules are requested via filter
	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("cloud_security_custom_rules") {
			for _, value := range filter.AcceptableValues {
				resources = append(resources, g.createResource(value))
			}
		}
	}

	if len(resources) > 0 {
		g.Resources = resources
		return nil
	}

	// Query all cloud security custom rules using the correct API
	// Filter for custom rules only (rule_origin=custom)
	filterStr := "rule_origin:'custom'"
	queryParams := &cloud_policies.QueryRuleParams{
		Context: context.Background(),
		Filter:  &filterStr,
	}

	resp, err := client.CloudPolicies.QueryRule(queryParams)
	if err != nil {
		return fmt.Errorf("failed to query cloud security custom rules: %v", err)
	}

	if resp.Payload == nil || resp.Payload.Resources == nil {
		// No custom rules found - this is valid, return empty list
		g.Resources = []terraformutils.Resource{}
		return nil
	}

	// Get detailed information for each rule
	if len(resp.Payload.Resources) > 0 {
		getParams := &cloud_policies.GetRuleParams{
			Context: context.Background(),
			Ids:     resp.Payload.Resources,
		}

		detailResp, err := client.CloudPolicies.GetRule(getParams)
		if err != nil {
			return fmt.Errorf("failed to get cloud security custom rule details: %v", err)
		}

		if detailResp.Payload != nil && detailResp.Payload.Resources != nil {
			// Convert API response to resource IDs
			ruleIDs := make([]string, 0, len(detailResp.Payload.Resources))
			for _, rule := range detailResp.Payload.Resources {
				if rule.UUID != nil {
					ruleIDs = append(ruleIDs, *rule.UUID)
				}
			}
			resources = g.createResources(ruleIDs)
		}
	}

	g.Resources = resources
	return nil
}

func (g *CloudSecurityCustomRuleGenerator) createResources(ruleIDs []string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, ruleID := range ruleIDs {
		if ruleID != "" {
			resources = append(resources, g.createResource(ruleID))
		}
	}
	return resources
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
	client := g.Args["client"].(*client.CrowdStrikeAPISpecification)

	// Check if specific frameworks are requested via filter
	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("cloud_compliance_frameworks") {
			for _, value := range filter.AcceptableValues {
				resources = append(resources, g.createResource(value))
			}
		}
	}

	if len(resources) > 0 {
		g.Resources = resources
		return nil
	}

	// Query all cloud compliance frameworks using the correct API
	// Filter for custom frameworks only (is_custom:true)
	filterStr := "is_custom:true"
	queryParams := &cloud_policies.QueryComplianceFrameworksParams{
		Context: context.Background(),
		Filter:  &filterStr,
	}

	resp, err := client.CloudPolicies.QueryComplianceFrameworks(queryParams)
	if err != nil {
		return fmt.Errorf("failed to query cloud compliance frameworks: %v", err)
	}

	if resp.Payload == nil || resp.Payload.Resources == nil {
		// No custom frameworks found - this is valid, return empty list
		g.Resources = []terraformutils.Resource{}
		return nil
	}

	// Get detailed information for each framework
	if len(resp.Payload.Resources) > 0 {
		getParams := &cloud_policies.GetComplianceFrameworksParams{
			Context: context.Background(),
			Ids:     resp.Payload.Resources,
		}

		detailResp, err := client.CloudPolicies.GetComplianceFrameworks(getParams)
		if err != nil {
			return fmt.Errorf("failed to get cloud compliance framework details: %v", err)
		}

		if detailResp.Payload != nil && detailResp.Payload.Resources != nil {
			// Convert API response to resource IDs - Resources contain UUID strings
			frameworkIDs := make([]string, 0, len(detailResp.Payload.Resources))
			for _, framework := range detailResp.Payload.Resources {
				if framework.UUID != "" {
					frameworkIDs = append(frameworkIDs, framework.UUID)
				}
			}
			resources = g.createResources(frameworkIDs)
		}
	}

	g.Resources = resources
	return nil
}

func (g *CloudComplianceFrameworkGenerator) createResources(frameworkIDs []string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, frameworkID := range frameworkIDs {
		if frameworkID != "" {
			resources = append(resources, g.createResource(frameworkID))
		}
	}
	return resources
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
	client := g.Args["client"].(*client.CrowdStrikeAPISpecification)

	// Check if specific groups are requested via filter
	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("cloud_groups") {
			for _, value := range filter.AcceptableValues {
				resources = append(resources, g.createResource(value))
			}
		}
	}

	if len(resources) > 0 {
		g.Resources = resources
		return nil
	}

	// Query all cloud groups using the correct API
	queryParams := &cloud_security.ListCloudGroupIDsExternalParams{
		Context: context.Background(),
	}

	resp, err := client.CloudSecurity.ListCloudGroupIDsExternal(queryParams)
	if err != nil {
		return fmt.Errorf("failed to query cloud groups: %v", err)
	}

	if resp.Payload == nil || resp.Payload.Resources == nil {
		// No cloud groups found - this is valid, return empty list
		g.Resources = []terraformutils.Resource{}
		return nil
	}

	// Get detailed information for each cloud group
	if len(resp.Payload.Resources) > 0 {
		getParams := &cloud_security.ListCloudGroupsByIDExternalParams{
			Context: context.Background(),
			Ids:     resp.Payload.Resources,
		}

		detailResp, err := client.CloudSecurity.ListCloudGroupsByIDExternal(getParams)
		if err != nil {
			return fmt.Errorf("failed to get cloud group details: %v", err)
		}

		if detailResp.Payload != nil && detailResp.Payload.Resources != nil {
			// Convert API response to resource IDs
			groupIDs := make([]string, 0, len(detailResp.Payload.Resources))
			for _, group := range detailResp.Payload.Resources {
				if group.ID != "" {
					groupIDs = append(groupIDs, group.ID)
				}
			}
			resources = g.createResources(groupIDs)
		}
	}

	g.Resources = resources
	return nil
}

func (g *CloudGroupGenerator) createResources(groupIDs []string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, groupID := range groupIDs {
		if groupID != "" {
			resources = append(resources, g.createResource(groupID))
		}
	}
	return resources
}
