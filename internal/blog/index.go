package blog

import (
	"github.com/marcjulianschwarz/go-blog/internal/blog/post"
	"github.com/marcjulianschwarz/go-blog/internal/blog/tag"
)

// An index is storing a list of posts and several utilities to easily
// access tags and posts
type Index struct {
	Posts      []*post.Post
	PostsById  map[string]*post.Post
	PostsByTag map[string][]*post.Post
	TagsById   map[string]*tag.Tag
}

func NewIndex() *Index {
	return &Index{
		PostsById:  make(map[string]*post.Post),
		PostsByTag: make(map[string][]*post.Post),
		TagsById:   make(map[string]*tag.Tag),
	}
}

func (i *Index) AddPost(post *post.Post) {
	i.Posts = append(i.Posts, post)
	i.PostsById[post.Id] = post

	for _, postTag := range post.Tags {
		i.TagsById[postTag.ID] = &postTag
		i.PostsByTag[postTag.ID] = append(i.PostsByTag[postTag.ID], post)
	}
}

func (i *Index) GetAllTags() []*tag.Tag {
	tags := make([]*tag.Tag, 0, len(i.TagsById))
	for _, tag := range i.TagsById {
		tags = append(tags, tag)
	}
	return tags
}
