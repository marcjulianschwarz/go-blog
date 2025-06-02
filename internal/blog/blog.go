package blog

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/marcjulianschwarz/go-blog/internal/config"
	"github.com/marcjulianschwarz/go-blog/internal/template"
	"github.com/marcjulianschwarz/go-blog/internal/yaml"
)

type BlogService struct {
	config          *config.BlogConfig
	templateService *template.TemplateService
	index           *Index
}

func NewBlogService(config *config.BlogConfig) *BlogService {

	return &BlogService{
		config:          config,
		index:           NewIndex(),
		templateService: template.NewTemplateService(config),
	}
}

// Reads all markdown files from the input path. Then extracts the YAML contained
// inside of the frontmatter. Adds a resulting post to the index or skips it when
// needed (e.g. skip attribute set or published date in the future or includes demo tag)
func (b *BlogService) ReadPosts() {
	fileSystem := os.DirFS(b.config.InputPath)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		fileExt := filepath.Ext(path)
		if fileExt == ".md" {
			data, readErr := fs.ReadFile(fileSystem, path)
			if readErr != nil {
				log.Fatal(readErr)
			}

			post := Post{}
			post.Content = string(data)

			yaml, err := yaml.GetPostYAML(post.Content)
			if err != nil {
				fmt.Println("could not get post YAML, skipping post", path)
				return nil
			}
			post.YAML = yaml

			if post.YAML.Skip {
				fmt.Printf("Skipping %s\n", post)
				return nil
			}

			fmt.Printf("Adding %s\n", post)
			b.index.AddPost(post)

			return nil
		}

		return nil
	})
}

// Creates an index.html file in the output path and writes the
// executed template filled with data from the current state of the index.
func (b *BlogService) WriteIndex() error {
	file, err := os.Create(filepath.Join(b.config.OutputPath, "index.html"))
	if err != nil {
		fmt.Println("could not write index")
		return err
	}

	postEntries := make([]template.PostEntryData, len(b.index.Posts))
	for i, post := range b.index.Posts {
		postEntries[i] = template.PostEntryData{
			URL:   "test url",
			Title: post.YAML.Title,
			Date:  post.YAML.Published,
		}
	}

	return b.templateService.Render(file, "index.html", template.TemplateData{
		PostList:          template.PostListData{Posts: postEntries},
		AllTagsList:       "all tags",
		RecentCount:       0,
		ArchivedPostsList: "all archived posts",
	})
}

func Main(config config.BlogConfig) {
	blogService := NewBlogService(&config)

	blogService.ReadPosts()
	err := blogService.WriteIndex()
	if err != nil {
		log.Fatal(err)
	}

}

// func X() {
// 	html := markdown.ToHTML([]byte(post.Content), nil, nil)
// 	filename := filepath.Base(path)
// 	htmlFilename := strings.Replace(filename, fileExt, "", -1) + ".html"

// 	htmlPath := filepath.Join(config.OutputPath, config.PostsSubPath, htmlFilename)
// 	errWrite := os.WriteFile(htmlPath, html, 0644)
// 	if errWrite != nil {
// 		log.Fatal("error writing")
// 	}
// }
