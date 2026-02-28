// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/reduce/vibedropper-cli/internal/mocktest"
)

func TestCustomersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"customers", "retrieve",
		"--customer-id", "customerId",
	)
}

func TestCustomersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"customers", "update",
		"--customer-id", "customerId",
		"--name", "name",
		"--pickup-location-id", "pickupLocationId",
		"--region-id", "regionId",
	)
}

func TestCustomersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"customers", "list",
		"--limit", "0",
		"--page", "0",
		"--search", "search",
	)
}
