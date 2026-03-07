// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/reduce/vibedropper-cli/internal/mocktest"
)

func TestCustomersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "customers", "retrieve",
			"--api-key", "string",
			"--customer-id", "customerId",
		)
	})
}

func TestCustomersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "customers", "update",
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"addressLine1: addressLine1\n" +
			"addressLine2: addressLine2\n" +
			"city: city\n" +
			"country: country\n" +
			"firstName: firstName\n" +
			"lastName: lastName\n" +
			"name: name\n" +
			"pickupLocationId: pickupLocationId\n" +
			"postalCode: postalCode\n" +
			"regionId: regionId\n" +
			"state: state\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "customers", "update",
			"--api-key", "string",
			"--customer-id", "customerId",
		)
	})
}

func TestCustomersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "customers", "list",
			"--api-key", "string",
			"--limit", "100",
			"--page", "0",
			"--search", "search",
		)
	})
}
