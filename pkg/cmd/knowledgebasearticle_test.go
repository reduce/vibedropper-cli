// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/vibedropper-cli/internal/mocktest"
)

func TestKnowledgeBasesArticlesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"knowledge-bases:articles", "create",
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
		"--kb-id", "kbId",
		"--limit", "100",
		"--page", "0",
	)
}
