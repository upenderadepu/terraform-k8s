package workspace

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testOrganization = "hashicorp-team-demo"
)

func setupClient(t *testing.T) *TerraformCloudClient {
	if os.Getenv("TF_ACC") == "" {
		t.Skipf("this test requires TF_CLI_CONFIG_FILE and TF_URL for Terraform Enterprise; set TF_ACC=1 to run it")
	}
	tfClient := &TerraformCloudClient{
		Organization: testOrganization,
	}
	err := tfClient.GetClient()
	assert.NoError(t, err)
	return tfClient
}

func TestOrganizationTerraformCloud(t *testing.T) {
	dir, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("Could not retrieve home directory for terraform cloud credentials, %s", err)
	}
	os.Setenv("TF_URL", "")
	os.Setenv("TF_CLI_CONFIG_FILE", fmt.Sprintf("%s/.terraformrc", dir))
	tfClient := setupClient(t)
	err = tfClient.CheckOrganization()
	assert.NoError(t, err)
}

func TestOrganizationTerraformEnterprise(t *testing.T) {
	tfClient := setupClient(t)
	err := tfClient.CheckOrganization()
	assert.NoError(t, err)
}

func TestOrganizationNotFound(t *testing.T) {
	tfClient := setupClient(t)
	tfClient.Organization = "doesnotexist"
	err := tfClient.CheckOrganization()
	assert.Error(t, err)
}
