package blog

import (
	"fmt"

	"github.com/marcjulianschwarz/go-blog/internal/yaml"
)

type Post struct {
	Content string
	YAML    yaml.PostYAML
}

func (p Post) String() string {
	return fmt.Sprintf("Post{%s}", p.YAML.Title)
}
