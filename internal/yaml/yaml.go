package yaml

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type PostYAML struct {
	Title     string   `yaml:"blog-title"`
	Subtitle  string   `yaml:"blog-subtitle"`
	Published string   `yaml:"blog-published"`
	Tags      []string `yaml:"blog-tags"`
	Archived  bool     `yaml:"blog-archived"`
	Skip      bool     `yaml:"blog-skip"`
	Author    string   `yaml:"blog-author"`
}

func getFrontmatter(content string) (frontmatter string, remainingContent string, found bool) {
	_, after, found := strings.Cut(content, "---")
	if !found {
		return "", "", false
	}
	frontmatter, remainingContent, found = strings.Cut(after, "---")
	if !found {
		return "", "", false
	}
	return strings.TrimSpace(frontmatter), strings.TrimSpace(remainingContent), true
}

func GetPostYAML(content string) (postYAMl PostYAML, frontmatter string, blogContent string, err error) {
	frontmatter, blogContent, found := getFrontmatter(content)
	if !found {
		fmt.Println("no frontmatter found")
	}

	postYAML := PostYAML{}
	err = yaml.Unmarshal([]byte(frontmatter), &postYAML)
	return postYAML, frontmatter, blogContent, err
}
