package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/Masterminds/sprig/v3"
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
	logger.Output(2, fmt.Sprintln(args...))
}

func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Output(2, fmt.Sprintln(args...))
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Output(2, fmt.Sprintln(args...))
}

//
// For HTTP error responses
//

// RenderableError wraps an error and adds information that can be
// rendered in a template to show the user.
type RenderableError struct {
	status  int
	userMsg string
	err     error
}

// Error returns the error as a readable string.
func (r *RenderableError) Error() string {
	return r.err.Error()
}

// Unwrap returns the error wrapped by RenderableError.
func (r *RenderableError) Unwrap() error {
	return r.err
}

func newNotFoundError(err error) error {
	return &RenderableError{
		status:  http.StatusNotFound,
		userMsg: "Not Found",
		err:     err,
	}
}

func newInternalServerError(err error) error {
	return &RenderableError{
		status:  http.StatusInternalServerError,
		userMsg: "Something went wrong",
		err:     err,
	}
}

func serveInvalidCsrfToken(w http.ResponseWriter, r *http.Request) {
	serveError(w, r, &RenderableError{
		status:  http.StatusForbidden,
		userMsg: "It's not possible to do this right now. Please go back, reload, and try again.",
		err:     errors.New("the user made a request with an invalid csrf token"),
	})
}

// serveError will render a templated error page. If the error is a RenderableError
// the userMsg stored in it will be displayed to the user.
func serveError(w http.ResponseWriter, r *http.Request, err error) {
	var rerr *RenderableError
	if !errors.As(err, &rerr) {
		serveError(w, r, newInternalServerError(fmt.Errorf("an unrenderable error ocurred: %v", err)))
		return
	}
	if rerr.status >= 500 && rerr.status < 600 {
		danger(rerr)
	}
	w.WriteHeader(rerr.status)
	if isAuthenticated(r) {
		generateHTML(w, r, "error", rerr.userMsg, "layout", "private.navbar", "error")
	} else {
		generateHTML(w, r, "error", rerr.userMsg, "layout", "public.navbar", "error")
	}
}

//
// Other
//

type viewData struct {
	PageName string
	Data     interface{}
}

func generateHTML(w http.ResponseWriter, r *http.Request, pageName string, data interface{}, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(
		template.New("layout").
			Funcs(sprig.FuncMap()).
			Funcs(templateFuncs(r)).
			ParseFiles(files...),
	)
	err := templates.Execute(w, viewData{
		PageName: pageName,
		Data:     data,
	})
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

// https://blog.questionable.services/article/generating-secure-random-numbers-crypto-rand/
// generateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// https://blog.questionable.services/article/generating-secure-random-numbers-crypto-rand/
// generateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomString(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
