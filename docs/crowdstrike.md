# Use Terraformer with [CrowdStrike Falcon](https://www.crowdstrike.com/)

This provider uses the [terraform-provider-crowdstrike](https://registry.terraform.io/providers/CrowdStrike/crowdstrike/latest).

## Usage

### 1. Installation
First you will need to install Terraformer with the CrowdStrike provider. See the [README](https://github.com/GoogleCloudPlatform/terraformer#installation).

### 2. Set up a template Terraform workspace
Before you can use Terraformer, you need to create a template workspace so that Terraformer
can access the [CrowdStrike/crowdstrike](https://registry.terraform.io/providers/CrowdStrike/crowdstrike/latest) provider.

To do this, create a new directory with a basic `provider.tf` file:
```hcl
terraform {
  required_providers {
    crowdstrike = {
      source  = "CrowdStrike/crowdstrike"
      version = "~> 0.5"
    }
  }
}

provider "crowdstrike" {
  # Configuration options
}
```

then run:
```bash
$ terraform init
```

You should see the output: `Terraform has been successfully initialized!`

### 3. Set up authentication

You can authenticate using environment variables or command-line flags:

#### Environment Variables
```bash
export FALCON_CLIENT_ID="your-client-id"
export FALCON_CLIENT_SECRET="your-client-secret"
export FALCON_CLOUD="us-1"  # Optional: us-1, us-2, eu-1, us-gov-1, or autodiscover
export FALCON_MEMBER_CID="member-cid"  # Optional: For MSSP scenarios
```

### 4. Run Terraformer

```bash
# Import all supported resources
terraformer import crowdstrike --resources="*"

# Import specific resources
terraformer import crowdstrike --resources="host_groups,prevention_policies"

# Use command-line flags for authentication
terraformer import crowdstrike --resources="host_groups" \
  --client-id="your-client-id" \
  --client-secret="your-client-secret" \
  --cloud="us-1"
```

### 5. Inspect the imported Terraform files

You should now see a `generated/` subdirectory with generated files.

You can now initialize and use your new generated resources:
```bash
$ cd generated/crowdstrike/
$ terraform init
$ terraform plan # No changes. Your infrastructure matches the configuration.
```

### Filtering Resources

You can use the `filter` argument to restrict the import of Terraform resources.

Filtering based on resource ID:

```bash
# Import specific host groups by ID
terraformer import crowdstrike --resources=host_groups --filter="host_groups=id1:id2:id3"

# Import specific prevention policies by ID
terraformer import crowdstrike --resources=prevention_policies --filter="prevention_policies=policy-id-1:policy-id-2"
```

### Path Output

You can specify a custom output directory:

```bash
terraformer import crowdstrike --resources="host_groups" --path-output="./crowdstrike-terraform"
```

## Supported CrowdStrike resources

*   `host_groups`
    * `crowdstrike_host_group`
*   `prevention_policies`
    * `crowdstrike_prevention_policy_windows`
    * `crowdstrike_prevention_policy_mac`
    * `crowdstrike_prevention_policy_linux`
*   `sensor_update_policies`
    * `crowdstrike_sensor_update_policy`
*   `content_update_policies`
    * `crowdstrike_content_update_policy`
*   `sensor_visibility_exclusions`
    * `crowdstrike_sensor_visibility_exclusion`
*   `fim_policies`
    * `crowdstrike_filevantage_policy`
*   `filevantage_rule_groups`
    * `crowdstrike_filevantage_rule_group`
*   `it_automation_tasks`
    * `crowdstrike_it_automation_task`
*   `it_automation_task_groups`
    * `crowdstrike_it_automation_task_group`
*   `it_automation_policies`
    * `crowdstrike_it_automation_policy`
*   `cloud_aws_accounts`
    * `crowdstrike_cloud_aws_account`
*   `cloud_azure_tenants`
    * `crowdstrike_cloud_azure_tenant`
*   `cloud_security_custom_rules`
    * `crowdstrike_cloud_security_custom_rule`
*   `cloud_compliance_frameworks`
    * `crowdstrike_cloud_compliance_custom_framework`
*   `cloud_groups`
    * `crowdstrike_cloud_group`

## API Scopes Required

Ensure your CrowdStrike API credentials have the following scopes for the resources you want to import:

- **Host groups**: Read
- **Prevention policies**: Read  
- **Sensor update policies**: Read
- **Content update policies**: Read
- **Sensor visibility exclusions**: Read
- **FileVantage (FIM)**: Read
- **IT Automation**: Read & Write
- **Cloud security**: Read & Write
- **CSPM registration**: Read & Write

## Notes

- **Sensitive Data**: API credentials (client_id, client_secret) are not stored in the generated Terraform files. You'll need to configure these separately in your provider block.
- **MSSP**: If you're using a master CID to access a member CID, set the `FALCON_MEMBER_CID` environment variable.
- **Cloud Regions**: The provider supports multiple Falcon cloud regions (us-1, us-2, eu-1, us-gov-1). Use `autodiscover` to automatically detect the correct region.

