package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
)

// templateFuncs returns a map of functions that can be used in the templates
func templateFuncs(r *http.Request) template.FuncMap {
	return template.FuncMap{
		"csrf":             tmplCsrfTag(r),
		"hasFlash":         tmplHasFlash(r),
		"flash":            tmplFlash(r),
		"bootrapAlertType": tmplBootstrapAlertType,
		"old":              tmplOld(r),
		"quantity":         tmplQuantity,
	}
}

// tmplCsrfTag generates a form input that holds the CSRF token
func tmplCsrfTag(r *http.Request) func() template.HTML {
	return func() template.HTML {
		return csrf.TemplateField(r)
	}
}

// tmplHasFlash checks if the current request contains a flash message
func tmplHasFlash(r *http.Request) func() bool {
	return func() bool {
		return hasFlash(r)
	}
}

// tmplFlash returns the flash message
func tmplFlash(r *http.Request) func() *flashMessage {
	return func() *flashMessage {
		return getFlash(r)
	}
}

// tmplBootstrapAlertType converts the flash message type
// to a bootstrap alert type
func tmplBootstrapAlertType(typ flashType) string {
	types := map[flashType]string{
		flashInfo:    "primary",
		flashSuccess: "success",
		flashWarning: "warning",
		flashDanger:  "danger",
	}
	if bt, ok := types[typ]; ok {
		return bt
	}
	panic("Cannot convert flashType to bootstrap alert type")
}

// tmplOld returns the previously submitted value of the
// form field, if available.
func tmplOld(r *http.Request) func(field string) string {
	return func(field string) string {
		return oldFormValue(r, field)
	}
}

// tmplQuantity displays the value shortened with metric suffix.
// The quantity can be hovered to see the complete value.
func tmplQuantity(value int64) template.HTML {
	full := fmtQuantityFull(value)
	short := fmtQuantityShort(value)
	return template.HTML(fmt.Sprintf("<span title=\"%s\">%s</span>", full, short))
}
