package cmd

import (
	"log"

	"github.com/marcjulianschwarz/go-blog/internal/blog"
)

func Generate(blogService *blog.BlogService) {

	blogService.ReadPosts()
	blogService.SortPosts()

	err := blogService.WriteIndex()
	if err != nil {
		log.Fatal("failed writing index", err)
	}

	err = blogService.WritePosts()
	if err != nil {
		log.Fatal("failed writing posts", err)
	}

	blogService.WriteTagPages()
	blogService.WriteSitemap()
}
