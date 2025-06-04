package sitemap

import (
	"encoding/xml"
	"os"
	"path/filepath"
)

type Sitemap struct {
	urls []SitemapURL
}

type SitemapURL struct {
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod"`
}

type SitemapXML struct {
	XMLName xml.Name     `xml:"urlset"`
	Xmlns   string       `xml:"xmlns,attr"`
	URLs    []SitemapURL `xml:"url"`
}

func NewSitemap() *Sitemap {
	return &Sitemap{
		urls: make([]SitemapURL, 0),
	}
}

func (s *Sitemap) UpdateSitemap(url, lastmod string) {
	s.urls = append(s.urls, SitemapURL{
		Loc: url, LastMod: lastmod,
	})
}

func (s *Sitemap) SaveSitemap(path string) error {
	sitemapXML := SitemapXML{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  s.urls,
	}

	file, err := os.Create(filepath.Join(path, "sitemap.xml"))
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(xml.Header)

	encoder := xml.NewEncoder(file)
	encoder.Indent("", "  ")

	return encoder.Encode(sitemapXML)
}
