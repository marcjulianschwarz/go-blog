package template

import (
	"html/template"
	"io"

	"github.com/marcjulianschwarz/go-blog/internal/blog/post"
	"github.com/marcjulianschwarz/go-blog/internal/config"
)

type PostEntryData struct {
	URL   string
	Title string
	Date  string
}

type PostData struct {
	Title    string
	Subtitle string
	Date     string
	Content  template.HTML
}

type PostListData struct {
	Posts []post.Post
}

type TagData struct {
	URL   string
	Color string
	Name  string
}

type TagPageData struct {
	Tag      TagData
	Count    uint
	PostList PostListData
}

type IndexData struct {
	RecentCount       uint
	PostList          PostListData
	Post              PostData
	ArchivedPostsList string
	AllTagsList       string
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

func (t *TemplateService) RenderIndex(wr io.Writer, data IndexData) error {
	return t.tmpl.ExecuteTemplate(wr, "index.html", data)
}

func (t *TemplateService) RenderPost(wr io.Writer, data PostData) error {
	return t.tmpl.ExecuteTemplate(wr, "post.html", data)
}

func (t *TemplateService) RenderTagPage(wr io.Writer, data TagPageData) error {
	return t.tmpl.ExecuteTemplate(wr, "tag-page.html", data)
}
