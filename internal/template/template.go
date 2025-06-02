package template

import (
	"html/template"
	"io"

	"github.com/marcjulianschwarz/go-blog/internal/config"
)

type PostEntryData struct {
	URL   string
	Title string
	Date  string
}

type PostListData struct {
	Posts []PostEntryData
}

type TemplateData struct {
	RecentCount       uint
	PostList          PostListData
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

func (t *TemplateService) Render(wr io.Writer, name string, data TemplateData) error {
	return t.tmpl.ExecuteTemplate(wr, name, data)
}
