package main

import (
	"html/template"
	"net/http"
)

type (
	templateData struct {
		Header headerData
		Body   any
	}

	headerData struct{}
)

var templates *template.Template

func renderTemplate(w http.ResponseWriter, template string, data any) error {
	passData := templateData{
		Header: headerData{},
		Body:   data,
	}

	return templates.ExecuteTemplate(w, template, passData)
}

func renderSecureTemplate(w http.ResponseWriter, template string, data any) {
	passData := templateData{
		Header: headerData{},
		Body:   data,
	}

	if err := templates.ExecuteTemplate(w, template, passData); err != nil {
		panic(err)
	}
}
