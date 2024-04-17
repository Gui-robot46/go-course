package render

import (
	"fmt"
	"net/http"
	"text/template"
)

const url string = "./templates/"

// This is a public function
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var url_full string = url + tmpl
	//fmt.Println(url_full)
	parsedTemplate, _ := template.ParseFiles(url_full)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template", err)
		return
	}
}
