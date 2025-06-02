package template

import (
	"html/template"
	"io"

	"github.com/marcjulianschwarz/go-blog/internal/config"
)

type TemplateData struct {
	RecentCount       uint
	AllTagsList       string
	ArchivedPostsList string
	Header            string
	AllPostsList      string
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

func (t *TemplateService) Render(wr io.Writer, name string, data TemplateData) {

	err := t.tmpl.ExecuteTemplate(wr, name, data)
	if err != nil {
		panic(err)
	}

}
