# CrowdStrike Provider for Terraformer

This provider allows you to import existing CrowdStrike Falcon resources into Terraform configuration files.

## Prerequisites

1. CrowdStrike Falcon API credentials (Client ID and Client Secret)
2. Appropriate API scopes for the resources you want to import

## Supported Resources

- `host_groups` - Host Groups
- `prevention_policies` - Prevention Policies (Windows, Mac, Linux)
- `sensor_update_policies` - Sensor Update Policies
- `content_update_policies` - Content Update Policies
- `sensor_visibility_exclusions` - Sensor Visibility Exclusions
- `fim_policies` - File Integrity Monitoring Policies (placeholder)
- `filevantage_rule_groups` - Filevantage Rule Groups (placeholder)
- `it_automation_tasks` - IT Automation Tasks (placeholder)
- `it_automation_task_groups` - IT Automation Task Groups (placeholder)
- `it_automation_policies` - IT Automation Policies (placeholder)
- `cloud_aws_accounts` - Cloud AWS Accounts (placeholder)
- `cloud_azure_tenants` - Cloud Azure Tenants (placeholder)
- `cloud_security_custom_rules` - Cloud Security Custom Rules (placeholder)
- `cloud_compliance_frameworks` - Cloud Compliance Frameworks (placeholder)
- `cloud_groups` - Cloud Groups (placeholder)

## Authentication

You can authenticate using environment variables or command-line flags:

### Environment Variables
```bash
export FALCON_CLIENT_ID="your-client-id"
export FALCON_CLIENT_SECRET="your-client-secret"
export FALCON_CLOUD="us-1"  # Optional: us-1, us-2, eu-1, us-gov-1, or autodiscover
export FALCON_MEMBER_CID="member-cid"  # Optional: For MSSP scenarios
```

### Command-line Flags
```bash
terraformer import crowdstrike --client-id="your-client-id" --client-secret="your-client-secret" --cloud="us-1"
```

## Usage Examples

### Import All Supported Resources
```bash
terraformer import crowdstrike --resources="*"
```

### Import Specific Resources
```bash
terraformer import crowdstrike --resources="host_groups,prevention_policies"
```

### Import with Filtering
```bash
terraformer import crowdstrike --resources="host_groups" --filter="host_groups=id1:id2:id3"
```

### Import to Specific Directory
```bash
terraformer import crowdstrike --resources="host_groups" --path-output="./crowdstrike-terraform"
```

### List Supported Resources
```bash
terraformer import crowdstrike list
```

## Output

The provider generates:
- `.tf` files with resource configurations
- `terraform.tfstate` file with the imported state
- `provider.tf` file with the CrowdStrike provider configuration

## Configuration

The generated `provider.tf` will look like:
```hcl
provider "crowdstrike" {
  client_id     = "your-client-id"
  client_secret = "your-client-secret"
  cloud         = "us-1"
  member_cid    = "member-cid"  # Optional
}
```

## Current Status

**Note**: This is the initial framework implementation. The core resource generators are implemented with basic structure, but the actual API calls to CrowdStrike Falcon APIs are not yet implemented. Currently, the generators return empty resource lists.

### Implemented
- ✅ Provider framework and CLI integration
- ✅ All resource generator structures
- ✅ Authentication handling
- ✅ Command-line interface
- ✅ Build system integration

### TODO
- 🔄 Implement actual gofalcon API calls
- 🔄 Add proper error handling and retry logic
- 🔄 Add resource filtering and pagination
- 🔄 Add comprehensive testing
- 🔄 Add resource relationships and dependencies

## Contributing

To implement the actual API calls:

1. Update the `InitResources()` method in each generator
2. Use the `client` from `g.Args["client"].(*client.CrowdStrikeAPISpecification)`
3. Call the appropriate gofalcon API methods
4. Parse the response and create terraform resources
5. Handle errors and edge cases

Example structure for implementing API calls:
```go
func (g *HostGroupGenerator) InitResources() error {
    client := g.Args["client"].(*client.CrowdStrikeAPISpecification)
    
    // Query host groups using gofalcon
    params := &host_group.QueryHostGroupsParams{
        Context: context.Background(),
    }
    
    resp, err := client.HostGroup.QueryHostGroups(params)
    if err != nil {
        return fmt.Errorf("failed to query host groups: %v", err)
    }
    
    // Process response and create resources
    g.Resources = g.createResources(resp.Payload.Resources)
    return nil
}
```

## API Scopes Required

Ensure your API credentials have the following scopes:
- Host groups: Read
- Prevention policies: Read  
- Sensor update policies: Read
- Content update policies: Read
- Sensor visibility exclusions: Read
- (Additional scopes for other resources as needed)
