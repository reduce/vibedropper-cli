// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/reduce/vibedropper-cli/internal/mocktest"
)

func TestListsSubscribersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"lists:subscribers", "list",
			"--list-id", "listId",
		)
	})
}

func TestListsSubscribersAdd(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"lists:subscribers", "add",
			"--list-id", "listId",
			"--email", "dev@stainless.com",
			"--custom-fields", "{}",
			"--name", "name",
			"--pickup-location-id", "pickupLocationId",
			"--region-id", "regionId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"email: dev@stainless.com\n" +
			"customFields: {}\n" +
			"name: name\n" +
			"pickupLocationId: pickupLocationId\n" +
			"regionId: regionId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"lists:subscribers", "add",
			"--list-id", "listId",
		)
	})
}

func TestListsSubscribersRemove(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"lists:subscribers", "remove",
			"--list-id", "listId",
			"--subscriber-id", "subscriberId",
		)
	})
}
