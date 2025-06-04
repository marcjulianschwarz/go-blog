package blog

import "github.com/marcjulianschwarz/go-blog/internal/blog/post"

type Index struct {
	Posts []post.Post
}

func NewIndex() *Index {
	return &Index{}
}

func (i *Index) AddPost(post post.Post) {
	i.Posts = append(i.Posts, post)
}

// Returns a list of unique tags that are currently in the index
func (i *Index) Tags() []string {
	tagSet := make(map[string]struct{})
	for _, post := range i.Posts {
		for _, tag := range post.YAML.Tags {
			tagSet[tag] = struct{}{}
		}
	}

	uniqueTags := make([]string, 0, len(tagSet))
	for tag := range tagSet {
		uniqueTags = append(uniqueTags, tag)
	}
	return uniqueTags
}
