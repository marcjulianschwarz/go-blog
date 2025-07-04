package blog

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/marcjulianschwarz/go-blog/internal/blog/post"
	"github.com/marcjulianschwarz/go-blog/internal/blog/tag"
	"github.com/marcjulianschwarz/go-blog/internal/config"
	"github.com/marcjulianschwarz/go-blog/internal/markdown"
	"github.com/marcjulianschwarz/go-blog/internal/sitemap"
	tpl "github.com/marcjulianschwarz/go-blog/internal/template"
	"github.com/marcjulianschwarz/go-blog/internal/yaml"
)

type BlogService struct {
	config          *config.BlogConfig
	templateService *tpl.TemplateService
	index           *Index
	sitemap         *sitemap.Sitemap
}

func NewBlogService(config *config.BlogConfig) *BlogService {

	return &BlogService{
		config:          config,
		templateService: tpl.NewTemplateService(config),
		index:           NewIndex(),
		sitemap:         sitemap.NewSitemap(),
	}
}

// WARNING: this deletes the entire output path
// only call this when you really want to delete the blog
func (b *BlogService) DeleteBlog() error {
	err := clearDirectory(b.config.OutputPath + "/" + b.config.PostsSubPath)
	err = clearDirectory(b.config.OutputPath + "/" + b.config.TagsSubPath)
	return err

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

			post := post.Post{}

			postYAML, _, blogContent, err := yaml.GetPostYAML(string(data))
			if err != nil {
				fmt.Println("could not get post YAML, skipping post", path)
				return nil
			}

			if postYAML.Skip {
				fmt.Printf("Skipping %s\n", filename)
				return nil
			}

			html := markdown.ToHTML(blogContent)
			html = strings.ReplaceAll(html, `src="/images/`, `src="/`+b.config.BlogSubPath+"/"+b.config.MediaSubPath+"/")

			post.Id = filename
			post.URL = "/" + b.config.BlogSubPath + "/" + b.config.PostsSubPath + "/" + post.Id + "/"
			post.Title = postYAML.Title
			post.Subtitle = postYAML.Subtitle
			post.Date = postYAML.Published
			post.Content = blogContent
			post.Author = postYAML.Author
			post.HTML = template.HTML(html)
			post.YAML = postYAML

			// create tags slice with capacity to hold the tags and an additional year
			tags := make([]tag.Tag, 0, len(postYAML.Tags)+1)
			for _, tagName := range postYAML.Tags {
				tagId := tag.TagNameToId(tagName)
				tags = append(tags, tag.Tag{
					Name:  tagName,
					URL:   "/" + b.config.BlogSubPath + "/" + b.config.TagsSubPath + "/" + tagId + "/",
					Color: "tag-post",
					ID:    tagId,
				})
			}

			year, _, found := strings.Cut(post.Date, "-")
			if found && len(year) == 4 {
				tags = append(tags, tag.Tag{
					Name:  year,
					URL:   "/" + b.config.BlogSubPath + "/" + b.config.TagsSubPath + "/" + year + "/",
					Color: "tag-year",
					ID:    year,
				})
			}

			post.Tags = tags

			// fmt.Printf("Adding %s\n", post)
			b.index.AddPost(&post)
			return nil
		}

		return nil
	})
}

func (b *BlogService) SortPosts() {
	b.index.SortByDate(true)
}

// Creates an index.html file in the output path and writes the
// executed template filled with data from the current state of the index.
func (b *BlogService) WriteIndex() error {
	file, err := os.Create(filepath.Join(b.config.OutputPath, "index.html"))
	if err != nil {
		fmt.Println("could not write index")
		return err
	}

	nonArchived := b.index.FilterNonArchived()
	return b.templateService.RenderIndex(file, tpl.IndexPageData{
		Posts:         nonArchived,
		Tags:          b.index.GetAllTags(),
		RecentCount:   0,
		ArchivedPosts: b.index.FilterArchived(),
		Meta: tpl.MetaData{
			Title:        "MJ's Blog",
			Description:  "MJ's Blog",
			Keywords:     "",
			CanonicalURL: b.config.PublishURL + b.config.BlogSubPath,
			Author:       "Marc Julian Schwarz",
		},
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

		err = b.templateService.RenderPost(file, tpl.PostPageData{
			Post: post,
			Meta: tpl.MetaData{
				Title:        post.Title,
				Description:  truncateWithMinMax(post.Content, 130, 160),
				Keywords:     tag.TagsToString(post.Tags),
				CanonicalURL: b.config.PublishURL + post.URL,
				Author:       post.Author,
			},
		})
	}
	return nil
}

// Creates tag pages for all tags by creating a directory with the
// tag's name and rendering the tag page into an index.html file
// inside of that directory
func (b *BlogService) WriteTagPages() error {

	for tagId, posts := range b.index.PostsByTag {
		path := filepath.Join(b.config.OutputPath, b.config.TagsSubPath, tagId)
		err := os.MkdirAll(path, 0755)
		if err != nil {
			fmt.Println("could not write tag directory", err)
			continue
		}

		file, err := os.Create(filepath.Join(path, "index.html"))
		if err != nil {
			fmt.Println("could not write tag page", err)
			continue
		}

		post.SortPostsByDate(posts, true)

		err = b.templateService.RenderTagPage(file, tpl.TagPageData{
			Tag:   *b.index.TagsById[tagId],
			Posts: posts,
			Count: len(posts),
		})

		if err != nil {
			fmt.Println("could not render tag page", err)
			continue
		}
	}
	return nil
}

func (b *BlogService) WriteSitemap() {
	for _, post := range b.index.Posts {
		b.sitemap.UpdateSitemap(b.config.PublishURL+post.URL, post.YAML.Published) // TODO: use last mod
	}

	for _, tag := range b.index.GetAllTags() {
		b.sitemap.UpdateSitemap(b.config.PublishURL+tag.URL, "")
	}

	b.sitemap.SaveSitemap(b.config.OutputPath)
}
