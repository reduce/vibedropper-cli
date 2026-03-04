// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/reduce/vibedropper-cli/internal/mocktest"
)

func TestCampaignsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"campaigns", "retrieve",
		"--api-key", "string",
		"--campaign-id", "campaignId",
	)
}

func TestCampaignsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"campaigns", "list",
		"--api-key", "string",
	)
}
