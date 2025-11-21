#!/bin/bash

# CrowdStrike Development Environment Setup
# This script helps you set up your local environment for testing

echo "🔧 CrowdStrike Development Environment Setup"
echo "============================================"

# Create a local environment file (not committed to git)
ENV_FILE=".env.crowdstrike.local"

echo ""
echo "This script will help you set up your CrowdStrike development credentials."
echo "Your credentials will be stored in: $ENV_FILE"
echo ""
echo "⚠️  IMPORTANT: This file will NOT be committed to git (it's in .gitignore)"
echo ""

# Check if env file already exists
if [ -f "$ENV_FILE" ]; then
    echo "📁 Found existing environment file: $ENV_FILE"
    echo ""
    read -p "Do you want to update it? (y/n): " update_env
    if [ "$update_env" != "y" ]; then
        echo "Using existing environment file."
        echo "To load it, run: source $ENV_FILE"
        exit 0
    fi
fi

echo "Please provide your CrowdStrike development environment details:"
echo ""

# Get Client ID
read -p "Falcon Client ID: " client_id
if [ -z "$client_id" ]; then
    echo "❌ Client ID is required"
    exit 1
fi

# Get Client Secret (hidden input)
echo -n "Falcon Client Secret: "
read -s client_secret
echo ""
if [ -z "$client_secret" ]; then
    echo "❌ Client Secret is required"
    exit 1
fi

# Get Cloud Region
echo ""
echo "Available cloud regions:"
echo "  1) autodiscover (recommended for dev)"
echo "  2) us-1"
echo "  3) us-2" 
echo "  4) eu-1"
echo "  5) us-gov-1"
echo ""
read -p "Select cloud region (1-5) [1]: " cloud_choice
cloud_choice=${cloud_choice:-1}

case $cloud_choice in
    1) cloud="autodiscover" ;;
    2) cloud="us-1" ;;
    3) cloud="us-2" ;;
    4) cloud="eu-1" ;;
    5) cloud="us-gov-1" ;;
    *) cloud="autodiscover" ;;
esac

# Get Member CID (optional for MSSP)
echo ""
read -p "Member CID (optional, for MSSP scenarios): " member_cid

# Create environment file
cat > "$ENV_FILE" << EOF
# CrowdStrike Development Environment Configuration
# Generated on $(date)
# 
# ⚠️  DO NOT COMMIT THIS FILE TO GIT!
#
# To use these credentials, run: source $ENV_FILE

export FALCON_CLIENT_ID="$client_id"
export FALCON_CLIENT_SECRET="$client_secret"
export FALCON_CLOUD="$cloud"
EOF

if [ -n "$member_cid" ]; then
    echo "export FALCON_MEMBER_CID=\"$member_cid\"" >> "$ENV_FILE"
fi

cat >> "$ENV_FILE" << EOF

# Convenience aliases for testing
alias tf-cs='./terraformer-crowdstrike'
alias tf-cs-list='./terraformer-crowdstrike import crowdstrike list'
alias tf-cs-test='./test_with_real_credentials.sh'

echo "🔐 CrowdStrike credentials loaded for development environment"
echo "Cloud: \$FALCON_CLOUD"
echo "Client ID: \${FALCON_CLIENT_ID:0:8}..."
EOF

chmod 600 "$ENV_FILE"

echo ""
echo "✅ Environment file created: $ENV_FILE"
echo ""
echo "🔒 File permissions set to 600 (owner read/write only)"
echo ""
echo "To load your credentials, run:"
echo "  source $ENV_FILE"
echo ""
echo "Then you can test with:"
echo "  ./test_with_real_credentials.sh"
echo ""
echo "Or run commands directly:"
echo "  ./terraformer-crowdstrike import crowdstrike list"
echo "  ./terraformer-crowdstrike import crowdstrike --resources=host_groups"
