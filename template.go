package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
)

// templateFuncs returns a map of functions that can be used in the templates
func templateFuncs(r *http.Request) template.FuncMap {
	return template.FuncMap{
		"csrf": tmplCsrfTag(r),
	}
}

// tmplCsrfTag generates a form input that holds the CSRF token
func tmplCsrfTag(r *http.Request) func() template.HTML {
	return func() template.HTML {
		return csrf.TemplateField(r)
	}
}
