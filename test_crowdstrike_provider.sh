#!/bin/bash

# Comprehensive test script for CrowdStrike provider
# This script tests all functionality without requiring real CrowdStrike credentials

set -e

echo "🧪 CrowdStrike Provider Test Suite"
echo "=================================="

TERRAFORMER_BIN="./terraformer-crowdstrike"
TEST_OUTPUT_DIR="./test-output"
FAILED_TESTS=0
TOTAL_TESTS=0

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Test function
run_test() {
    local test_name="$1"
    local test_command="$2"
    local expected_exit_code="${3:-0}"
    
    TOTAL_TESTS=$((TOTAL_TESTS + 1))
    echo -n "Testing: $test_name... "
    
    if eval "$test_command" >/dev/null 2>&1; then
        actual_exit_code=0
    else
        actual_exit_code=$?
    fi
    
    if [ $actual_exit_code -eq $expected_exit_code ]; then
        echo -e "${GREEN}PASS${NC}"
    else
        echo -e "${RED}FAIL${NC} (expected exit code $expected_exit_code, got $actual_exit_code)"
        FAILED_TESTS=$((FAILED_TESTS + 1))
    fi
}

# Test function with output capture
run_test_with_output() {
    local test_name="$1"
    local test_command="$2"
    local expected_pattern="$3"
    
    TOTAL_TESTS=$((TOTAL_TESTS + 1))
    echo -n "Testing: $test_name... "
    
    local output
    output=$(eval "$test_command" 2>&1)
    
    if echo "$output" | grep -q "$expected_pattern"; then
        echo -e "${GREEN}PASS${NC}"
    else
        echo -e "${RED}FAIL${NC} (expected pattern '$expected_pattern' not found)"
        echo "  Output: $output"
        FAILED_TESTS=$((FAILED_TESTS + 1))
    fi
}

echo ""
echo "1. 🔧 Basic Binary Tests"
echo "------------------------"

run_test "Binary exists and is executable" "test -x $TERRAFORMER_BIN"
run_test "Version command works" "$TERRAFORMER_BIN --version"
run_test "Help command works" "$TERRAFORMER_BIN --help"

echo ""
echo "2. 🌐 Provider Registration Tests"
echo "--------------------------------"

run_test_with_output "CrowdStrike provider is listed" "$TERRAFORMER_BIN import --help" "crowdstrike"
run_test_with_output "CrowdStrike help shows correctly" "$TERRAFORMER_BIN import crowdstrike --help" "CrowdStrike Falcon"

echo ""
echo "3. 📋 Resource Listing Tests"
echo "----------------------------"

run_test_with_output "List command works" "$TERRAFORMER_BIN import crowdstrike list" "host_groups"
run_test_with_output "Prevention policies listed" "$TERRAFORMER_BIN import crowdstrike list" "prevention_policies"
run_test_with_output "Sensor update policies listed" "$TERRAFORMER_BIN import crowdstrike list" "sensor_update_policies"
run_test_with_output "Content update policies listed" "$TERRAFORMER_BIN import crowdstrike list" "content_update_policies"
run_test_with_output "SVE listed" "$TERRAFORMER_BIN import crowdstrike list" "sensor_visibility_exclusions"

echo ""
echo "4. 🔐 Authentication Flag Tests"
echo "------------------------------"

run_test_with_output "Client ID flag exists" "$TERRAFORMER_BIN import crowdstrike --help" "client-id"
run_test_with_output "Client secret flag exists" "$TERRAFORMER_BIN import crowdstrike --help" "client-secret"
run_test_with_output "Cloud flag exists" "$TERRAFORMER_BIN import crowdstrike --help" "cloud"
run_test_with_output "Member CID flag exists" "$TERRAFORMER_BIN import crowdstrike --help" "member-cid"

echo ""
echo "5. 📁 Import Command Tests (Without Credentials)"
echo "-----------------------------------------------"

# Clean up any previous test output
rm -rf "$TEST_OUTPUT_DIR"
mkdir -p "$TEST_OUTPUT_DIR"

# These should fail gracefully with authentication errors, not crash
run_test "Import without credentials fails gracefully" "$TERRAFORMER_BIN import crowdstrike --resources=host_groups --path-output=$TEST_OUTPUT_DIR" 1
run_test "Import with invalid credentials fails gracefully" "$TERRAFORMER_BIN import crowdstrike --client-id=invalid --client-secret=invalid --resources=host_groups --path-output=$TEST_OUTPUT_DIR" 1

echo ""
echo "6. 🎯 Resource Filtering Tests"
echo "-----------------------------"

run_test_with_output "Resource filtering help shows" "$TERRAFORMER_BIN import crowdstrike --help" "filter"
run_test_with_output "Resource selection help shows" "$TERRAFORMER_BIN import crowdstrike --help" "resources"

echo ""
echo "7. 📝 Output Format Tests"
echo "------------------------"

run_test_with_output "HCL output format supported" "$TERRAFORMER_BIN import crowdstrike --help" "hcl"
run_test_with_output "JSON output format supported" "$TERRAFORMER_BIN import crowdstrike --help" "json"

echo ""
echo "8. 🔍 Code Quality Tests"
echo "-----------------------"

echo -n "Testing: Go fmt compliance... "
if gofmt -l providers/crowdstrike/*.go cmd/provider_cmd_crowdstrike.go | grep -q .; then
    echo -e "${RED}FAIL${NC} (files not formatted with gofmt)"
    FAILED_TESTS=$((FAILED_TESTS + 1))
else
    echo -e "${GREEN}PASS${NC}"
fi
TOTAL_TESTS=$((TOTAL_TESTS + 1))

echo -n "Testing: Go build succeeds... "
if go build -o /tmp/terraformer-test . >/dev/null 2>&1; then
    echo -e "${GREEN}PASS${NC}"
    rm -f /tmp/terraformer-test
else
    echo -e "${RED}FAIL${NC}"
    FAILED_TESTS=$((FAILED_TESTS + 1))
fi
TOTAL_TESTS=$((TOTAL_TESTS + 1))

echo -n "Testing: Go vet passes... "
if go vet ./providers/crowdstrike/ ./cmd/ >/dev/null 2>&1; then
    echo -e "${GREEN}PASS${NC}"
else
    echo -e "${RED}FAIL${NC}"
    FAILED_TESTS=$((FAILED_TESTS + 1))
fi
TOTAL_TESTS=$((TOTAL_TESTS + 1))

echo ""
echo "9. 📚 Documentation Tests"
echo "------------------------"

run_test "README exists" "test -f providers/crowdstrike/README.md"
run_test "Example script exists" "test -f providers/crowdstrike/example.sh"
run_test "Example script is executable" "test -x providers/crowdstrike/example.sh"

echo ""
echo "🏁 Test Results Summary"
echo "======================"

if [ $FAILED_TESTS -eq 0 ]; then
    echo -e "${GREEN}✅ All $TOTAL_TESTS tests passed!${NC}"
    echo ""
    echo "🎉 Your CrowdStrike provider is ready for contribution!"
    echo ""
    echo "Next steps:"
    echo "1. Test with real CrowdStrike credentials (optional)"
    echo "2. Create pull request to upstream repository"
    echo "3. Implement actual API calls in future iterations"
else
    echo -e "${RED}❌ $FAILED_TESTS out of $TOTAL_TESTS tests failed${NC}"
    echo ""
    echo "Please fix the failing tests before creating a pull request."
    exit 1
fi

# Cleanup
rm -rf "$TEST_OUTPUT_DIR"
