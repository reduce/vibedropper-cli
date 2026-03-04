// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/reduce/vibedropper-cli/internal/mocktest"
)

func TestKnowledgeBasesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"knowledge-bases", "retrieve",
		"--kb-id", "kbId",
	)
}

func TestKnowledgeBasesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"knowledge-bases", "update",
		"--kb-id", "kbId",
		"--description", "description",
		"--name", "name",
		"--sort-order", "0",
	)
}

func TestKnowledgeBasesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"knowledge-bases", "list",
	)
}

func TestKnowledgeBasesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"knowledge-bases", "delete",
		"--kb-id", "kbId",
	)
}
