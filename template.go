package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/csrf"
	"github.com/pdeguing/empire-and-foundation/data"
	"github.com/pdeguing/empire-and-foundation/ent/timer"
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
		"duration":         tmplDuration,
		"byTimerGroup":     tmplByTimerGroup,
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

// tmplDuration display the duration in a way that the front-end
// can update the value as time progresses without having to reload
// the page.
func tmplDuration(d time.Duration) template.HTML {
	dInt := int64(d.Seconds())
	atTime := time.Now().Add(d).Format("02-01-2006 15:04:05 MST")
	dStr := fmt.Sprint(d.Round(time.Second))
	return template.HTML(fmt.Sprintf("<span title=\"%s\" data-duration=\"%d\">%s</span>", atTime, dInt, dStr))
}

// tmplByTimerGroup returns the timer in group g from map m or nil if it
// is does not exist in the map.
func tmplByTimerGroup(m map[timer.Group]*data.Timer, g string) *data.Timer {
	return m[timer.Group(g)]
}
