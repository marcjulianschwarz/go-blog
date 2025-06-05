package main

import (
	"log"

	"github.com/marcjulianschwarz/go-blog/internal/blog"
	"github.com/marcjulianschwarz/go-blog/internal/config"
)

func main() {

	config, err := config.Get("./")
	if err != nil {
		log.Fatal("no config file found")
	}

	blog.Main(config)
}
