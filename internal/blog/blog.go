package blog

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/marcjulianschwarz/go-blog/internal/config"
	tpl "github.com/marcjulianschwarz/go-blog/internal/template"
	"github.com/marcjulianschwarz/go-blog/internal/yaml"
)

type BlogService struct {
	config          *config.BlogConfig
	templateService *tpl.TemplateService
	index           *Index
}

func NewBlogService(config *config.BlogConfig) *BlogService {

	return &BlogService{
		config:          config,
		index:           NewIndex(),
		templateService: tpl.NewTemplateService(config),
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

			filename := filepath.Base(path)
			filename = strings.ReplaceAll(filename, fileExt, "")

			post := Post{}
			post.Id = filename
			post.Content = string(data)
			post.URL = b.config.PublishPath + "/" + b.config.PostsSubPath + "/" + post.Id

			postYAML, _, blogContent, err := yaml.GetPostYAML(post.Content)
			if err != nil {
				fmt.Println("could not get post YAML, skipping post", path)
				return nil
			}
			post.YAML = postYAML
			post.HTML = string(markdown.ToHTML([]byte(blogContent), nil, nil))

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

	postEntries := make([]tpl.PostEntryData, len(b.index.Posts))
	for i, post := range b.index.Posts {
		postEntries[i] = tpl.PostEntryData{
			URL:   post.URL,
			Title: post.YAML.Title,
			Date:  post.YAML.Published,
		}
	}

	return b.templateService.RenderIndex(file, tpl.IndexData{
		PostList:          tpl.PostListData{Posts: postEntries},
		AllTagsList:       "all tags",
		RecentCount:       0,
		ArchivedPostsList: "all archived posts",
	})
}

func (b *BlogService) WritePosts() error {
	for _, post := range b.index.Posts {
		path := filepath.Join(b.config.OutputPath, b.config.PostsSubPath, post.Id)
		err := os.MkdirAll(path, 0755)
		if err != nil {
			fmt.Println("could not write post directory", err)
			continue
		}

		file, err := os.Create(filepath.Join(path, "index.html"))
		if err != nil {
			fmt.Println("could not write post")
			continue
		}

		err = b.templateService.RenderPost(file, tpl.PostData{
			Title:    post.YAML.Title,
			Subtitle: post.YAML.Subtitle,
			Date:     post.YAML.Published,
			Content:  template.HTML(post.HTML),
		})
	}
	return nil
}

func Main(config config.BlogConfig) {
	blogService := NewBlogService(&config)

	blogService.ReadPosts()
	err := blogService.WriteIndex()
	if err != nil {
		log.Fatal(err)
	}

	err = blogService.WritePosts()
	if err != nil {
		log.Fatal(err)
	}

}
