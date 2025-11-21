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
	"errors"
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/crowdstrike/gofalcon/falcon"
	"github.com/crowdstrike/gofalcon/falcon/client"
	"github.com/zclconf/go-cty/cty"
)

type CrowdStrikeProvider struct { //nolint
	terraformutils.Provider
	clientID     string
	clientSecret string
	cloud        string
	memberCID    string
	client       *client.CrowdStrikeAPISpecification
}

// Init check env params and initialize Falcon Client
func (p *CrowdStrikeProvider) Init(args []string) error {
	// Parse arguments: [clientID, clientSecret, cloud, memberCID]
	if len(args) >= 1 && args[0] != "" {
		p.clientID = args[0]
	} else if clientID := os.Getenv("FALCON_CLIENT_ID"); clientID != "" {
		p.clientID = clientID
	} else {
		return errors.New("client-id requirement: use --client-id or set FALCON_CLIENT_ID environment variable")
	}

	if len(args) >= 2 && args[1] != "" {
		p.clientSecret = args[1]
	} else if clientSecret := os.Getenv("FALCON_CLIENT_SECRET"); clientSecret != "" {
		p.clientSecret = clientSecret
	} else {
		return errors.New("client-secret requirement: use --client-secret or set FALCON_CLIENT_SECRET environment variable")
	}

	if len(args) >= 3 && args[2] != "" {
		p.cloud = args[2]
	} else if cloud := os.Getenv("FALCON_CLOUD"); cloud != "" {
		p.cloud = cloud
	} else {
		p.cloud = "autodiscover" // Default value
	}

	if len(args) >= 4 && args[3] != "" {
		p.memberCID = args[3]
	} else if memberCID := os.Getenv("FALCON_MEMBER_CID"); memberCID != "" {
		p.memberCID = memberCID
	}

	// Initialize the CrowdStrike Falcon client
	apiConfig := falcon.ApiConfig{
		Cloud:        falcon.Cloud(p.cloud),
		ClientId:     p.clientID,
		ClientSecret: p.clientSecret,
		Context:      context.Background(),
	}

	if p.memberCID != "" {
		apiConfig.MemberCID = p.memberCID
	}

	client, err := falcon.NewClient(&apiConfig)
	if err != nil {
		return fmt.Errorf("failed to create CrowdStrike Falcon client: %v", err)
	}

	p.client = client
	return nil
}

// GetName return string of provider name for CrowdStrike
func (p *CrowdStrikeProvider) GetName() string {
	return "crowdstrike"
}

// GetConfig return map of provider config for CrowdStrike
func (p *CrowdStrikeProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"client_id":     cty.StringVal(p.clientID),
		"client_secret": cty.StringVal(p.clientSecret),
		"cloud":         cty.StringVal(p.cloud),
		"member_cid":    cty.StringVal(p.memberCID),
	})
}

// InitService ...
func (p *CrowdStrikeProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"client":        p.client,
		"client_id":     p.clientID,
		"client_secret": p.clientSecret,
		"cloud":         p.cloud,
		"member_cid":    p.memberCID,
	})
	return nil
}

// GetSupportedService return map of support service for CrowdStrike
func (p *CrowdStrikeProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		// Core resources
		"host_groups":                  &HostGroupGenerator{},
		"prevention_policies":          &PreventionPolicyGenerator{},
		"sensor_update_policies":       &SensorUpdatePolicyGenerator{},
		"content_update_policies":      &ContentUpdatePolicyGenerator{},
		"sensor_visibility_exclusions": &SensorVisibilityExclusionGenerator{},

		// FIM resources
		"fim_policies":            &FIMPolicyGenerator{},
		"filevantage_rule_groups": &FilevantageRuleGroupGenerator{},

		// IT Automation resources
		"it_automation_tasks":       &ITAutomationTaskGenerator{},
		"it_automation_task_groups": &ITAutomationTaskGroupGenerator{},
		"it_automation_policies":    &ITAutomationPolicyGenerator{},

		// Cloud resources
		"cloud_aws_accounts":          &CloudAWSAccountGenerator{},
		"cloud_azure_tenants":         &CloudAzureTenantGenerator{},
		"cloud_security_custom_rules": &CloudSecurityCustomRuleGenerator{},
		"cloud_compliance_frameworks": &CloudComplianceFrameworkGenerator{},
		"cloud_groups":                &CloudGroupGenerator{},
	}
}

// GetResourceConnections return map of resource connections for CrowdStrike
func (p CrowdStrikeProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		"prevention_policies": {
			"host_groups": {
				"host_group_ids", "id",
			},
		},
		"sensor_update_policies": {
			"host_groups": {
				"host_group_ids", "id",
			},
		},
		"content_update_policies": {
			"host_groups": {
				"host_group_ids", "id",
			},
		},
		"fim_policies": {
			"host_groups": {
				"host_group_ids", "id",
			},
			"filevantage_rule_groups": {
				"rule_group_ids", "id",
			},
		},
		"it_automation_policies": {
			"host_groups": {
				"host_group_ids", "id",
			},
			"it_automation_task_groups": {
				"task_group_ids", "id",
			},
		},
		"it_automation_task_groups": {
			"it_automation_tasks": {
				"task_ids", "id",
			},
		},
	}
}

// GetProviderData return map of provider data for CrowdStrike
func (p CrowdStrikeProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}
