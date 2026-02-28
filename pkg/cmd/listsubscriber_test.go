// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/reduce/vibedropper-cli/internal/mocktest"
)

func TestListsSubscribersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"lists:subscribers", "list",
		"--list-id", "listId",
	)
}

func TestListsSubscribersAdd(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"lists:subscribers", "add",
		"--list-id", "listId",
		"--email", "dev@stainless.com",
		"--custom-fields", "{}",
		"--name", "name",
		"--pickup-location-id", "pickupLocationId",
		"--region-id", "regionId",
	)
}

func TestListsSubscribersRemove(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"lists:subscribers", "remove",
		"--list-id", "listId",
		"--subscriber-id", "subscriberId",
	)
}
