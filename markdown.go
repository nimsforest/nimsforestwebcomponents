package nwc

import (
	"bytes"
	"html/template"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

var markdownRenderer = goldmark.New(goldmark.WithExtensions(extension.Table))

// RenderMarkdown converts markdown source to HTML using goldmark.
// Returns the source as-is if rendering fails.
func RenderMarkdown(src string) template.HTML {
	var buf bytes.Buffer
	if err := markdownRenderer.Convert([]byte(src), &buf); err != nil {
		return template.HTML(template.HTMLEscapeString(src))
	}
	return template.HTML(buf.String())
}

// MarkdownEditorData is the data contract for the shared markdown editor templates.
type MarkdownEditorData struct {
	Content     string // Raw markdown content
	SaveURL     string // Form POST target URL
	CancelURL   string // Cancel link href
	Target      string // HTMX swap target (e.g. "#content-panel")
	Filename    string // Optional: filename for new documents
	IsNew       bool   // True when creating, false when editing
	Placeholder string // Textarea placeholder text
	Rows        int    // Textarea rows (0 defaults to 24)
}

// EditorRows returns the configured row count or 24 as default.
func (d MarkdownEditorData) EditorRows() int {
	if d.Rows > 0 {
		return d.Rows
	}
	return 24
}
