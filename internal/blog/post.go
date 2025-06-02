package blog

import (
	"fmt"

	"github.com/marcjulianschwarz/go-blog/internal/yaml"
)

type Post struct {
	Id      string
	Content string
	HTML    string
	URL     string
	YAML    yaml.PostYAML
}

func (p Post) String() string {
	return fmt.Sprintf("Post{%s}", p.YAML.Title)
}
