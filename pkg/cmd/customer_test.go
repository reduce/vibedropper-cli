// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/vibedropper-cli/internal/mocktest"
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
		"--address-line1", "addressLine1",
		"--address-line2", "addressLine2",
		"--city", "city",
		"--country", "country",
		"--first-name", "firstName",
		"--last-name", "lastName",
		"--name", "name",
		"--pickup-location-id", "pickupLocationId",
		"--postal-code", "postalCode",
		"--region-id", "regionId",
		"--state", "state",
	)
}

func TestCustomersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"customers", "list",
		"--limit", "100",
		"--page", "0",
		"--search", "search",
	)
}
