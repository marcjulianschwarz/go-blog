# Go Blog

## Configure Blog

Create a file called `config.yaml`. Example:

```yaml
input-path: "path/to/markdown"
output-path: "path/to/blog"
tags-subpath: "tags"
posts-subpath: "posts"
media-subpath: "media"
templates-path: "path/to/templates"
publish-url: "https://example.com/blog"
```

## Generate Blog

```bash
go run .
```
