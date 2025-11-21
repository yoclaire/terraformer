#!/bin/bash

# CrowdStrike Terraformer Example Script
# This script demonstrates how to use the CrowdStrike provider with Terraformer

set -e

echo "CrowdStrike Terraformer Example"
echo "==============================="

# Check if required environment variables are set
if [ -z "$FALCON_CLIENT_ID" ] || [ -z "$FALCON_CLIENT_SECRET" ]; then
    echo "Error: Please set FALCON_CLIENT_ID and FALCON_CLIENT_SECRET environment variables"
    echo ""
    echo "Example:"
    echo "export FALCON_CLIENT_ID=\"your-client-id\""
    echo "export FALCON_CLIENT_SECRET=\"your-client-secret\""
    echo "export FALCON_CLOUD=\"us-1\"  # Optional"
    echo ""
    exit 1
fi

# Set default cloud if not specified
if [ -z "$FALCON_CLOUD" ]; then
    export FALCON_CLOUD="autodiscover"
fi

echo "Using CrowdStrike Cloud: $FALCON_CLOUD"
echo ""

# Create output directory
OUTPUT_DIR="./crowdstrike-terraform"
mkdir -p "$OUTPUT_DIR"

echo "1. Listing supported resources..."
terraformer import crowdstrike list

echo ""
echo "2. Importing Host Groups..."
terraformer import crowdstrike \
    --resources="host_groups" \
    --path-output="$OUTPUT_DIR" \
    --verbose

echo ""
echo "3. Importing Prevention Policies..."
terraformer import crowdstrike \
    --resources="prevention_policies" \
    --path-output="$OUTPUT_DIR" \
    --verbose

echo ""
echo "4. Importing Sensor Update Policies..."
terraformer import crowdstrike \
    --resources="sensor_update_policies" \
    --path-output="$OUTPUT_DIR" \
    --verbose

echo ""
echo "5. Generated files:"
find "$OUTPUT_DIR" -name "*.tf" -o -name "*.tfstate" | head -10

echo ""
echo "Example complete! Check the $OUTPUT_DIR directory for generated Terraform files."
echo ""
echo "Note: The current implementation returns empty resource lists as the actual"
echo "CrowdStrike API calls are not yet implemented. This is the framework structure."
