package markdown

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html" // ← Aliased
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html" // ← Add this import
	"github.com/gomarkdown/markdown/parser"
)

type ChromaRenderer struct {
	*html.Renderer
}

func (r *ChromaRenderer) RenderNode(w io.Writer, node ast.Node, entering bool) ast.WalkStatus {
	if codeBlock, ok := node.(*ast.CodeBlock); ok && entering {
		lang := strings.TrimSpace(string(codeBlock.Info))
		code := string(codeBlock.Literal)

		lexer := lexers.Get(lang)
		if lexer == nil {
			lexer = lexers.Fallback
		}

		style := styles.Get("github")
		formatter := chromahtml.New(chromahtml.WithClasses(true))

		iterator, err := lexer.Tokenise(nil, code)
		if err != nil {
			fmt.Printf("ERROR: Tokenisation failed: %v\n", err)
		} else {
			var buf bytes.Buffer
			err = formatter.Format(&buf, style, iterator)
			if err != nil {
				fmt.Printf("ERROR: Formatting failed: %v\n", err)
			} else {
				w.Write(buf.Bytes())
				return ast.GoToNext
			}
		}
	}

	return r.Renderer.RenderNode(w, node, entering)
}

func ToHTML(content string) string {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	renderer := &ChromaRenderer{
		Renderer: html.NewRenderer(html.RendererOptions{
			Flags: htmlFlags,
		}),
	}

	doc := p.Parse([]byte(content))
	return string(markdown.Render(doc, renderer))
}
