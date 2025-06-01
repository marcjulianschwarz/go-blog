package yaml

import (
	"testing"
)

func TestGetFrontmatter(t *testing.T) {
	content := "---\ntest1\ntest2\n---\nmore text"
	frontmatter, remainingContent, found := getFrontmatter(content)

	if !found {
		t.Error("no frontmatter found")
	}
	if !(frontmatter == "test1\ntest2") {
		t.Error("frontmatter wrong")
	}
	if !(remainingContent == "more text") {
		t.Error("remaining content wrong")
	}
}
