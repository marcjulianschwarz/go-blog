package cmd

import (
	"log"

	"github.com/marcjulianschwarz/go-blog/internal/blog"
)

func Delete(blogService *blog.BlogService) error {
	println("Deleting Blog")
	err := blogService.DeleteBlog()
	if err != nil {
		log.Fatal("failed deleting blog", err)
	}
	return err
}
