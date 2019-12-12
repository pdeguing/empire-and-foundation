package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var logger *log.Logger

func init() {
	logFile, err := os.OpenFile("empire-and-foundation.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	logger = log.New(mw, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

//
// For logging
//

func info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

//
// For HTTP error responses
//

func internalServerError(w http.ResponseWriter, r *http.Request, err error, internalError string) {
	danger(err, internalError)
	respondWithError(w, r, "Something went wrong on our end.", http.StatusInternalServerError)
}

func invalidCsrfToken(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, r, "It's not possible to do this right now. Please go back, reload, and try again.", 403)
}

// respondWithError will render a templated error page. userMsg will
// be shown as a message to the user.
func respondWithError(w http.ResponseWriter, r *http.Request, userMsg string, code int) {
	w.WriteHeader(code)
	if isAuthenticated(r) {
		generateHTML(w, r, userMsg, "layout", "private.navbar", "error")
	} else {
		generateHTML(w, r, userMsg, "layout", "public.navbar", "error")
	}
}

//
// Other
//

func generateHTML(w http.ResponseWriter, r *http.Request, data interface{}, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.New("layout").Funcs(templateFuncs(r)).ParseFiles(files...))
	err := templates.Execute(w, data)
	if err != nil {
		danger(err, "unable to render template")
		// Executing the template doesn't even work so respond
		// with the most basic error message as a fallback.
		http.Error(w, "Something went wrong on our end.", 500)
	}
}

// fmtQuantityFull returns the quantity value as a formatted string.
func fmtQuantityFull(value int64) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%d", value)
}

// fmtQuantityShort returns value as a rounded number with metric suffix.
// The jump to the next suffix is deliberately done a little late. For
// example, if you have just 1800 metal, you would probably care to
// see when you reach 2000 metal. Once the value is over 5000 the jump
// is made and it will just show 5k. For other suffixes the same
// mechanism is used.
func fmtQuantityShort(value int64) string {
	p := message.NewPrinter(language.English)
	switch {
	case value >= 10e15:
		return p.Sprintf("%dP", value/1e15)
	case value >= 10e12:
		return p.Sprintf("%dT", value/1e12)
	case value >= 10e9:
		return p.Sprintf("%dG", value/1e9)
	case value >= 10e6:
		return p.Sprintf("%dM", value/1e6)
	case value >= 10e3:
		return p.Sprintf("%dk", value/1e3)
	default:
		return p.Sprintf("%d", value)
	}
}
