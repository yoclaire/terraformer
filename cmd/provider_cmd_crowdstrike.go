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

package cmd

import (
	crowdstrike_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/crowdstrike"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdCrowdStrikeImporter(options ImportOptions) *cobra.Command {
	var clientID, clientSecret, cloud, memberCID string
	cmd := &cobra.Command{
		Use:   "crowdstrike",
		Short: "Import current state to Terraform configuration from CrowdStrike Falcon",
		Long:  "Import current state to Terraform configuration from CrowdStrike Falcon",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newCrowdStrikeProvider()
			err := Import(provider, options, []string{clientID, clientSecret, cloud, memberCID})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newCrowdStrikeProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options,
		"host_groups,prevention_policies,sensor_update_policies,content_update_policies,sensor_visibility_exclusions",
		"host_groups=id1:id2:id4")
	cmd.PersistentFlags().StringVarP(&clientID, "client-id", "", "", "CrowdStrike Falcon Client ID or env param FALCON_CLIENT_ID")
	cmd.PersistentFlags().StringVarP(&clientSecret, "client-secret", "", "", "CrowdStrike Falcon Client Secret or env param FALCON_CLIENT_SECRET")
	cmd.PersistentFlags().StringVarP(&cloud, "cloud", "", "", "CrowdStrike Falcon Cloud (autodiscover, us-1, us-2, eu-1, us-gov-1) or env param FALCON_CLOUD")
	cmd.PersistentFlags().StringVarP(&memberCID, "member-cid", "", "", "CrowdStrike Member CID for MSSP scenarios or env param FALCON_MEMBER_CID")
	return cmd
}

func newCrowdStrikeProvider() terraformutils.ProviderGenerator {
	return &crowdstrike_terraforming.CrowdStrikeProvider{}
}
