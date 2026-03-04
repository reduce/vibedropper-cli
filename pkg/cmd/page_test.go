// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/reduce/vibedropper-cli/internal/mocktest"
)

func TestPagesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"pages", "retrieve",
		"--page-id", "pageId",
	)
}

func TestPagesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"pages", "update",
		"--page-id", "pageId",
		"--description", "description",
		"--name", "name",
		"--status", "DRAFT",
	)
}

func TestPagesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"pages", "list",
		"--limit", "100",
		"--page", "0",
		"--status", "DRAFT",
	)
}

func TestPagesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"pages", "delete",
		"--page-id", "pageId",
	)
}
