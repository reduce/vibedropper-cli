// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/reduce/vibedropper-cli/internal/mocktest"
)

func TestKnowledgeBasesArticlesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"knowledge-bases:articles", "create",
			"--kb-id", "kbId",
			"--title", "title",
			"--category-id", "categoryId",
			"--content", "content",
			"--excerpt", "excerpt",
			"--published=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"title: title\n" +
			"categoryId: categoryId\n" +
			"content: content\n" +
			"excerpt: excerpt\n" +
			"published: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"knowledge-bases:articles", "create",
			"--kb-id", "kbId",
		)
	})
}

func TestKnowledgeBasesArticlesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"knowledge-bases:articles", "list",
			"--kb-id", "kbId",
			"--limit", "100",
			"--page", "0",
		)
	})
}
