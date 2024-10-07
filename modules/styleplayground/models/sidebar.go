package models

import "github.com/a-h/templ"

// Sidebar is a struct that represents the sidebar content
type Sidebar struct {
	// Content is a map of string to string that represents the sidebar content
	Title    string
	Content  templ.Component
	Path     string
	Children []Sidebar
}
