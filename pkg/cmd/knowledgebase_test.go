// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/reduce/vibedropper-cli/internal/mocktest"
)

func TestKnowledgeBasesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "knowledge-bases", "retrieve",
			"--api-key", "string",
			"--kb-id", "kbId",
		)
	})
}

func TestKnowledgeBasesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "knowledge-bases", "update",
			"--api-key", "string",
			"--kb-id", "kbId",
			"--description", "description",
			"--name", "name",
			"--sort-order", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: description\n" +
			"name: name\n" +
			"sortOrder: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "knowledge-bases", "update",
			"--api-key", "string",
			"--kb-id", "kbId",
		)
	})
}

func TestKnowledgeBasesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "knowledge-bases", "list",
			"--api-key", "string",
		)
	})
}

func TestKnowledgeBasesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "knowledge-bases", "delete",
			"--api-key", "string",
			"--kb-id", "kbId",
		)
	})
}
