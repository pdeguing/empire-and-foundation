package main

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/pdeguing/empire-and-foundation/data"
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

func internalServerError(w http.ResponseWriter, r *http.Request, err error, internalError string) {
	danger(err, internalError)
	respondWithError(w, r, "Something went wrong on our end.", http.StatusInternalServerError)
}

func respondWithError(w http.ResponseWriter, r *http.Request, userMsg string, code int) {
	w.WriteHeader(code)
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, userMsg, "layout", "public.navbar", "error")
	} else {
		generateHTML(w, userMsg, "layout", "private.navbar", "error")
	}
}

func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

func parseTemplatesFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

// for logging
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
