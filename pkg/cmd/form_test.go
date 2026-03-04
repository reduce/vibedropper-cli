// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/vibedropper-cli/internal/mocktest"
)

func TestFormsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"forms", "retrieve",
		"--form-id", "formId",
	)
}

func TestFormsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"forms", "update",
		"--form-id", "formId",
		"--description", "description",
		"--list-id", "listId",
		"--status", "DRAFT",
		"--success-message", "successMessage",
		"--title", "title",
	)
}

func TestFormsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"forms", "list",
		"--limit", "100",
		"--page", "0",
		"--status", "DRAFT",
	)
}

func TestFormsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"forms", "delete",
		"--form-id", "formId",
	)
}

func TestFormsListSubmissions(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"forms", "list-submissions",
		"--form-id", "formId",
		"--limit", "100",
		"--page", "0",
	)
}
