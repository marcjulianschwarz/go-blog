package post

import (
	"fmt"
	"html/template"

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
	YAML     yaml.PostYAML
	Tags     []tag.Tag
}

func (p Post) String() string {
	return fmt.Sprintf("Post{%s}", p.YAML.Title)
}
