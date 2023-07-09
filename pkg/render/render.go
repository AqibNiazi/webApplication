package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// renderTemplate uses html template
func RenderTemplate(w http.ResponseWriter, html string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + html)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error occured while parsing", err)
		return
	}

}
