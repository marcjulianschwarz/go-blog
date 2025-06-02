package config

import (
	"io/fs"
	"os"

	"gopkg.in/yaml.v3"
)

type BlogConfig struct {
	InputPath    string `yaml:"input-path"`
	OutputPath   string `yaml:"output-path"`
	TagsSubPath  string `yaml:"tags-subpath"`
	PostsSubPath string `yaml:"posts-subpath"`
	MediaSubPath string `yaml:"media-subpath"`
	TemplatePath string `yaml:"templates-path"`
	PublishPath  string `yaml:"publish-path"`
}

// Get BlogConfig from config.yaml file in path directory
func Get(path string) (config BlogConfig, error error) {
	filesystem := os.DirFS(path)
	content, err := fs.ReadFile(filesystem, "config.yaml")
	if err != nil {
		return BlogConfig{}, err
	}

	config = BlogConfig{}
	err = yaml.Unmarshal([]byte(content), &config)
	if err != nil {
		return BlogConfig{}, err
	}
	return config, nil
}
