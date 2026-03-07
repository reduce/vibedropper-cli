// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/reduce/vibedropper-cli/internal/mocktest"
)

func TestPagesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "pages", "retrieve",
			"--api-key", "string",
			"--page-id", "pageId",
		)
	})
}

func TestPagesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "pages", "update",
			"--api-key", "string",
			"--page-id", "pageId",
			"--description", "description",
			"--name", "name",
			"--status", "DRAFT",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: description\n" +
			"name: name\n" +
			"status: DRAFT\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "pages", "update",
			"--api-key", "string",
			"--page-id", "pageId",
		)
	})
}

func TestPagesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "pages", "list",
			"--api-key", "string",
			"--limit", "100",
			"--page", "0",
			"--status", "DRAFT",
		)
	})
}

func TestPagesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "pages", "delete",
			"--api-key", "string",
			"--page-id", "pageId",
		)
	})
}
