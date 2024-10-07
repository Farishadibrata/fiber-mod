package utils

import (
	"farishadibrata.com/fiber-modular/modules/styleplayground/models"
	"github.com/a-h/templ"
)

func ToMapStringInterface(sidebar []models.Sidebar) map[string]templ.Component {
	m := make(map[string]templ.Component)
	for _, v := range sidebar {
		m[v.Path] = v.Content
		// Flatten the children recursively
		if len(v.Children) > 0 {
			for k, v := range ToMapStringInterface(v.Children) {
				m[k] = v
			}
		}
	}
	return m
}
