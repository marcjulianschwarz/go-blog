package main

import (
	"fmt"
	"log"
	"os"

	"github.com/marcjulianschwarz/go-blog/cmd"
	"github.com/marcjulianschwarz/go-blog/internal/blog"
	"github.com/marcjulianschwarz/go-blog/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: blog <command>")
		fmt.Println("Commands:")
		fmt.Println("  generate - Generate blog")
		fmt.Println("  delete   - Delete all generated files")
		os.Exit(1)
	}

	config, err := config.Get("./")
	if err != nil {
		log.Fatal("no config file found")
	}

	blogService := blog.NewBlogService(&config)

	command := os.Args[1]

	switch command {
	case "generate":
		cmd.Generate(blogService)
	case "delete":
		cmd.Delete(blogService)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
