package template

import (
	"html/template"
	"io"

	"github.com/marcjulianschwarz/go-blog/internal/blog/post"
	"github.com/marcjulianschwarz/go-blog/internal/blog/tag"
	"github.com/marcjulianschwarz/go-blog/internal/config"
)

type MetaData struct {
	Title        string
	Description  string
	Keywords     string
	Author       string
	CanonicalURL string
}

type TagPageData struct {
	Tag   tag.Tag
	Count int
	Posts []*post.Post
}

type PostPageData struct {
	Post *post.Post
	Meta MetaData
}

type IndexPageData struct {
	RecentCount   uint
	Posts         []*post.Post
	Tags          []*tag.Tag
	ArchivedPosts []*post.Post
	Meta          MetaData
}

type TemplateService struct {
	config *config.BlogConfig
	tmpl   *template.Template
}

// Creates a new template service for the templates specified in the configuration
func NewTemplateService(config *config.BlogConfig) *TemplateService {
	tmpl, err := template.ParseGlob(config.TemplatePath + "/*.html")
	if err != nil {
		panic(err)
	}
	return &TemplateService{
		config: config,
		tmpl:   tmpl,
	}
}

func (t *TemplateService) RenderIndex(wr io.Writer, data IndexPageData) error {
	return t.tmpl.ExecuteTemplate(wr, "index.html", data)
}

func (t *TemplateService) RenderPost(wr io.Writer, data PostPageData) error {
	return t.tmpl.ExecuteTemplate(wr, "post.html", data)
}

func (t *TemplateService) RenderTagPage(wr io.Writer, data TagPageData) error {
	return t.tmpl.ExecuteTemplate(wr, "tag-page.html", data)
}
