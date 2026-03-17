// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/reduce/vibedropper-cli/internal/mocktest"
)

func TestFormsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"forms", "retrieve",
			"--form-id", "formId",
		)
	})
}

func TestFormsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"forms", "update",
			"--form-id", "formId",
			"--description", "description",
			"--list-id", "listId",
			"--status", "DRAFT",
			"--success-message", "successMessage",
			"--title", "title",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: description\n" +
			"listId: listId\n" +
			"status: DRAFT\n" +
			"successMessage: successMessage\n" +
			"title: title\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"forms", "update",
			"--form-id", "formId",
		)
	})
}

func TestFormsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"forms", "list",
			"--limit", "100",
			"--page", "0",
			"--status", "DRAFT",
		)
	})
}

func TestFormsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"forms", "delete",
			"--form-id", "formId",
		)
	})
}

func TestFormsListSubmissions(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"forms", "list-submissions",
			"--form-id", "formId",
			"--limit", "100",
			"--page", "0",
		)
	})
}
