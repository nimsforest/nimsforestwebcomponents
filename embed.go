package nwc

import (
	"embed"
	"io/fs"
)

//go:embed static/* templates/*
var templates embed.FS

// TemplateFS returns the embedded filesystem containing shared templates
// (layout.html, components.html). Use this when composing custom layouts
// that need access to nwc's shared template blocks.
func TemplateFS() fs.FS {
	return templates
}
