// Package nwc provides HTML layout, template helpers, and Tailwind CSS
// configuration for NimsForest web UIs. Forest branding: dark background,
// lush green, solar gold, vibrant life.
package nwc

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"strings"
	"time"
)

// NavItem represents a single navigation link in the app header.
type NavItem struct {
	Label string
	Href  string
}

// ChatWidgetConfig configures the embedded nim chat widget.
type ChatWidgetConfig struct {
	BaseURL    string // widget server URL, e.g. "https://chatwidget.nimsforest.mynimsforest.com"
	DefaultNim string // optional default nim selection
}

// AppConfig configures the shared layout for a specific NimsForest application.
type AppConfig struct {
	Name       string            // displayed after "Nims" in header, e.g. "Organize", "Forest"
	Emoji      string            // header emoji, e.g. "🌿", "🌲"
	NavItems   []NavItem         // app-specific navigation links
	Footer     string            // footer text (defaults to "NimsForest" if empty)
	ChatWidget *ChatWidgetConfig // nil = no chat widget
}

// Renderer composes shared layout templates with project-specific page templates.
type Renderer struct {
	templates map[string]*template.Template
	funcMap   template.FuncMap
	config    AppConfig
}

// PageData is the standard data envelope passed to all templates.
type PageData struct {
	Title string
	Data  any
	App   AppConfig
}

// FuncMap returns the standard template function map for NimsForest UIs.
func FuncMap() template.FuncMap {
	return template.FuncMap{
		"formatBytes": FormatBytes,
		"formatMB": func(bytes uint64) uint64 {
			return bytes / 1024 / 1024
		},
		"formatGB": func(bytes uint64) string {
			gb := float64(bytes) / (1024 * 1024 * 1024)
			if gb >= 1 {
				return template.HTMLEscapeString(formatFloat(gb) + " GB")
			}
			mb := float64(bytes) / (1024 * 1024)
			return template.HTMLEscapeString(formatFloat(mb) + " MB")
		},
		"formatDate": func(t time.Time) string {
			if t.IsZero() {
				return ""
			}
			return t.Format("02 Jan 2006")
		},
		"timeAgo": func(t time.Time) string {
			if t.IsZero() {
				return "never"
			}
			d := time.Since(t)
			switch {
			case d < time.Minute:
				return "just now"
			case d < time.Hour:
				m := int(d.Minutes())
				if m == 1 {
					return "1 minute ago"
				}
				return formatInt(m) + " minutes ago"
			case d < 24*time.Hour:
				h := int(d.Hours())
				if h == 1 {
					return "1 hour ago"
				}
				return formatInt(h) + " hours ago"
			default:
				days := int(d.Hours() / 24)
				if days == 1 {
					return "1 day ago"
				}
				return formatInt(days) + " days ago"
			}
		},
		"percentage": func(used, total uint64) float64 {
			if total == 0 {
				return 0
			}
			return float64(used) / float64(total) * 100
		},
		"statusColor": func(status string) string {
			switch status {
			case "running":
				return "text-forest-green"
			case "stopped":
				return "text-red-400"
			default:
				return "text-gray-400"
			}
		},
		"formatTimestamp": func(t time.Time) string {
			if t.IsZero() {
				return ""
			}
			return t.UTC().Format("2006-01-02 15:04:05 UTC")
		},
		"lower": strings.ToLower,
		"add":   func(a, b int) int { return a + b },
		"dict": func(pairs ...any) map[string]any {
			m := make(map[string]any, len(pairs)/2)
			for i := 0; i < len(pairs)-1; i += 2 {
				if k, ok := pairs[i].(string); ok {
					m[k] = pairs[i+1]
				}
			}
			return m
		},
	}
}

// FormatBytes formats bytes into a human-readable string.
func FormatBytes(bytes uint64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
		TB = GB * 1024
	)
	switch {
	case bytes >= TB:
		return formatFloat(float64(bytes)/float64(TB)) + " TB"
	case bytes >= GB:
		return formatFloat(float64(bytes)/float64(GB)) + " GB"
	case bytes >= MB:
		return formatFloat(float64(bytes)/float64(MB)) + " MB"
	case bytes >= KB:
		return formatFloat(float64(bytes)/float64(KB)) + " KB"
	default:
		return formatInt(int(bytes)) + " B"
	}
}

func formatFloat(f float64) string {
	if f == float64(int(f)) {
		return formatInt(int(f))
	}
	// One decimal place
	i := int(f * 10)
	whole := i / 10
	frac := i % 10
	if frac < 0 {
		frac = -frac
	}
	return formatInt(whole) + "." + formatInt(frac)
}

func formatInt(i int) string {
	s := ""
	if i == 0 {
		return "0"
	}
	neg := i < 0
	if neg {
		i = -i
	}
	for i > 0 {
		s = string(rune('0'+i%10)) + s
		i /= 10
	}
	if neg {
		s = "-" + s
	}
	return s
}

// NewRenderer creates a Renderer that combines the shared nwc layout with
// project-specific page templates from the given filesystem.
//
// projectFS should contain templates in the given dir (e.g. "templates").
// pages lists the template filenames to compose with the layout.
// extraFuncs is merged with the standard FuncMap.
// config provides app-specific branding and navigation for the shared layout.
func NewRenderer(projectFS fs.FS, dir string, pages []string, extraFuncs template.FuncMap, config AppConfig) *Renderer {
	fm := FuncMap()
	for k, v := range extraFuncs {
		fm[k] = v
	}

	if config.Footer == "" {
		config.Footer = "NimsForest"
	}

	r := &Renderer{
		templates: make(map[string]*template.Template),
		funcMap:   fm,
		config:    config,
	}

	// Read shared templates
	layoutBytes, err := fs.ReadFile(templates, "templates/layout.html")
	if err != nil {
		log.Fatalf("nwc: failed to read layout.html: %v", err)
	}
	componentsBytes, err := fs.ReadFile(templates, "templates/components.html")
	if err != nil {
		log.Fatalf("nwc: failed to read components.html: %v", err)
	}

	return r.buildTemplates(projectFS, dir, pages, nil, layoutBytes, componentsBytes, fm)
}

// NewRendererWithShared works like NewRenderer but also includes project-level
// shared templates (e.g. sidebar.html) that are parsed into each page template set.
func NewRendererWithShared(projectFS fs.FS, dir string, pages []string, shared []string, extraFuncs template.FuncMap, config AppConfig) *Renderer {
	fm := FuncMap()
	for k, v := range extraFuncs {
		fm[k] = v
	}

	if config.Footer == "" {
		config.Footer = "NimsForest"
	}

	r := &Renderer{
		templates: make(map[string]*template.Template),
		funcMap:   fm,
		config:    config,
	}

	layoutBytes, err := fs.ReadFile(templates, "templates/layout.html")
	if err != nil {
		log.Fatalf("nwc: failed to read layout.html: %v", err)
	}
	componentsBytes, err := fs.ReadFile(templates, "templates/components.html")
	if err != nil {
		log.Fatalf("nwc: failed to read components.html: %v", err)
	}

	// Read project-level shared templates
	var sharedBytes [][]byte
	for _, s := range shared {
		b, err := fs.ReadFile(projectFS, dir+"/"+s)
		if err != nil {
			log.Fatalf("nwc: failed to read shared template %s: %v", s, err)
		}
		sharedBytes = append(sharedBytes, b)
	}

	return r.buildTemplates(projectFS, dir, pages, sharedBytes, layoutBytes, componentsBytes, fm)
}

func (r *Renderer) buildTemplates(projectFS fs.FS, dir string, pages []string, sharedBytes [][]byte, layoutBytes, componentsBytes []byte, fm template.FuncMap) *Renderer {
	for _, page := range pages {
		path := dir + "/" + page
		pageBytes, err := fs.ReadFile(projectFS, path)
		if err != nil {
			log.Fatalf("nwc: failed to read %s: %v", path, err)
		}

		t := template.New("layout.html").Funcs(fm)
		template.Must(t.Parse(string(layoutBytes)))
		template.Must(t.New("components.html").Parse(string(componentsBytes)))
		for i, sb := range sharedBytes {
			template.Must(t.New(fmt.Sprintf("shared_%d", i)).Parse(string(sb)))
		}
		template.Must(t.New(page).Parse(string(pageBytes)))

		r.templates[page] = t
	}

	return r
}

// Render executes the named page template with data wrapped in PageData.
func (r *Renderer) Render(w http.ResponseWriter, page string, title string, data any) {
	t, ok := r.templates[page]
	if !ok {
		http.Error(w, "template not found: "+page, http.StatusInternalServerError)
		return
	}

	pd := PageData{
		Title: title,
		Data:  data,
		App:   r.config,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := t.ExecuteTemplate(w, "layout.html", pd); err != nil {
		log.Printf("nwc: render error for %s: %v", page, err)
	}
}

// RenderFragment executes only the "content" block of the named page template,
// without the surrounding layout. Used for HTMX partial responses.
func (r *Renderer) RenderFragment(w http.ResponseWriter, page string, title string, data any) {
	t, ok := r.templates[page]
	if !ok {
		http.Error(w, "template not found: "+page, http.StatusInternalServerError)
		return
	}

	pd := PageData{
		Title: title,
		Data:  data,
		App:   r.config,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := t.ExecuteTemplate(w, "content", pd); err != nil {
		log.Printf("nwc: render fragment error for %s: %v", page, err)
	}
}

// StaticHandler returns an http.Handler serving the embedded static files.
func StaticHandler() http.Handler {
	sub, err := fs.Sub(templates, "static")
	if err != nil {
		log.Fatalf("nwc: failed to create static sub-fs: %v", err)
	}
	return http.FileServer(http.FS(sub))
}
