package post

import (
	"fmt"
	"html/template"
	"sort"

	"github.com/marcjulianschwarz/go-blog/internal/blog/tag"
	"github.com/marcjulianschwarz/go-blog/internal/yaml"
)

type Post struct {
	Id       string
	Content  string
	HTML     template.HTML
	URL      string
	Title    string
	Subtitle string
	Date     string
	Author   string
	YAML     yaml.PostYAML
	Tags     []tag.Tag
}

func (p Post) String() string {
	return fmt.Sprintf("Post{%s}", p.YAML.Title)
}

func SortPostsByDate(posts []*Post, descending bool) {
	sort.Slice(posts, func(idx, jdx int) bool {
		if descending {
			return posts[idx].Date > posts[jdx].Date
		}
		return posts[idx].Date < posts[jdx].Date
	})
}
