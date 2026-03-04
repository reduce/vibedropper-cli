// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/reduce/vibedropper-cli/internal/mocktest"
)

func TestKnowledgeBasesArticlesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"knowledge-bases:articles", "create",
		"--api-key", "string",
		"--kb-id", "kbId",
		"--title", "title",
		"--category-id", "categoryId",
		"--content", "content",
		"--excerpt", "excerpt",
		"--published=true",
	)
}

func TestKnowledgeBasesArticlesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"knowledge-bases:articles", "list",
		"--api-key", "string",
		"--kb-id", "kbId",
		"--limit", "100",
		"--page", "0",
	)
}
