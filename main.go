package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/marcjulianschwarz/go-blog/internal/config"
	"github.com/marcjulianschwarz/go-blog/internal/yaml"
)

func main() {

	config, err := config.Get("./")
	if err != nil {
		log.Fatal("no config file found")
	}

	fileSystem := os.DirFS(config.InputPath)

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

			content := string(data)
			postYAML, err := yaml.GetPostYAML(content)
			if err != nil {
				fmt.Println("could not get post YAML, skipping post", path)
				return nil
			}

			if postYAML.Skip {
				fmt.Println("skip post", path)
				return nil
			}

			html := markdown.ToHTML([]byte(content), nil, nil)
			filename := filepath.Base(path)
			htmlFilename := strings.Replace(filename, fileExt, "", -1) + ".html"

			htmlPath := filepath.Join(config.OutputPath, config.PostsSubPath, htmlFilename)
			errWrite := os.WriteFile(htmlPath, html, 0644)
			if errWrite != nil {
				log.Fatal("error writing")
			}
			return nil
		}

		return nil
	})

}
