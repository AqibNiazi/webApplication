package config

import (
	"html/template"
)

// AppConfig holds the application configuration
type Appconfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
