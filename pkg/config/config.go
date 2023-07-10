package config

import (
	"html/template"
)

// AppConfig holds the application configuration
type Appconfig struct {
	TemplateCache map[string]*template.Template
}
