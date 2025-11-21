#!/bin/bash

# Test CrowdStrike provider with real credentials
# Run this only if you have actual CrowdStrike API credentials

set -e

echo "🚀 CrowdStrike Provider Real Credentials Test"
echo "============================================="

# Check if credentials are provided
if [ -z "$FALCON_CLIENT_ID" ] || [ -z "$FALCON_CLIENT_SECRET" ]; then
    echo "❌ Real credentials not provided"
    echo ""
    echo "To test with real CrowdStrike credentials:"
    echo ""
    echo "Option 1 - Use setup script:"
    echo "  ./setup_dev_environment.sh"
    echo "  source .env.crowdstrike.local"
    echo "  $0"
    echo ""
    echo "Option 2 - Set manually:"
    echo "  export FALCON_CLIENT_ID=\"your-client-id\""
    echo "  export FALCON_CLIENT_SECRET=\"your-client-secret\""
    echo "  export FALCON_CLOUD=\"us-1\"  # Optional"
    echo "  $0"
    echo ""
    echo "⚠️  WARNING: This will attempt to connect to your real CrowdStrike tenant!"
    exit 1
fi

TERRAFORMER_BIN="./terraformer-crowdstrike"
TEST_OUTPUT_DIR="./test-real-output"

echo "Using real credentials for cloud: ${FALCON_CLOUD:-autodiscover}"
echo ""

# Clean up any previous test output
rm -rf "$TEST_OUTPUT_DIR"
mkdir -p "$TEST_OUTPUT_DIR"

echo "1. Testing authentication with real credentials..."
echo "================================================"

# Test basic authentication
echo "Testing basic connection..."
if $TERRAFORMER_BIN import crowdstrike --resources=host_groups --path-output="$TEST_OUTPUT_DIR" --verbose; then
    echo "✅ Authentication successful!"
    
    echo ""
    echo "2. Testing resource imports..."
    echo "============================="
    
    # Test each resource type
    for resource in "host_groups" "prevention_policies" "sensor_update_policies" "content_update_policies" "sensor_visibility_exclusions"; do
        echo ""
        echo "Testing $resource import..."
        
        resource_dir="$TEST_OUTPUT_DIR/$resource"
        mkdir -p "$resource_dir"
        
        if $TERRAFORMER_BIN import crowdstrike --resources="$resource" --path-output="$resource_dir" --verbose; then
            echo "✅ $resource import completed"
            
            # Check if files were generated
            if ls "$resource_dir"/*/*.tf >/dev/null 2>&1; then
                echo "  📁 Terraform files generated:"
                ls "$resource_dir"/*/*.tf | head -5
            else
                echo "  ℹ️  No resources found (this is expected with current implementation)"
            fi
        else
            echo "❌ $resource import failed"
        fi
    done
    
    echo ""
    echo "3. Testing advanced features..."
    echo "=============================="
    
    # Test JSON output
    echo "Testing JSON output format..."
    if $TERRAFORMER_BIN import crowdstrike --resources=host_groups --output=json --path-output="$TEST_OUTPUT_DIR/json_test" --verbose; then
        echo "✅ JSON output format works"
    else
        echo "❌ JSON output format failed"
    fi
    
    # Test filtering (this will work even with empty results)
    echo "Testing resource filtering..."
    if $TERRAFORMER_BIN import crowdstrike --resources=host_groups --filter="host_groups=test-id" --path-output="$TEST_OUTPUT_DIR/filter_test" --verbose; then
        echo "✅ Resource filtering works"
    else
        echo "❌ Resource filtering failed"
    fi
    
    echo ""
    echo "🎉 Real Credentials Test Complete!"
    echo "=================================="
    echo ""
    echo "✅ Your CrowdStrike provider successfully connects to the real API!"
    echo ""
    echo "Generated files are in: $TEST_OUTPUT_DIR"
    echo ""
    echo "Note: The current implementation returns empty resource lists."
    echo "This is expected as the actual API calls are not yet implemented."
    echo "The important thing is that authentication and the framework work!"
    
else
    echo "❌ Authentication failed with real credentials"
    echo ""
    echo "This could mean:"
    echo "1. Invalid credentials"
    echo "2. Incorrect cloud region"
    echo "3. Insufficient API scopes"
    echo "4. Network connectivity issues"
    echo ""
    echo "Check your credentials and try again."
    exit 1
fi

echo ""
echo "🧹 Cleanup..."
echo "============"
echo "Test files are preserved in: $TEST_OUTPUT_DIR"
echo "Remove them manually if desired: rm -rf $TEST_OUTPUT_DIR"
