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

	// templateService := template.NewTemplateService(&config)

	// file, err := os.Create(filepath.Join(config.OutputPath, "index.html"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close() // close file after function execution

	// data := template.TemplateData{
	// 	RecentCount:       3,
	// 	AllTagsList:       "all tags",
	// 	ArchivedPostsList: "archived list",
	// 	Header:            "some header",
	// 	AllPostsList:      "posts",
	// }

	// templateService.Render(file, "index.html", data)

}
