package main

import (
	"bytes"
	"errors"
	"fmt"
	"gopkg.in/gomail.v2"
	"html/template"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"

	"github.com/pdeguing/empire-and-foundation/ent"
)

var sendEmail = sendEmailSmtp

func sendEmailSmtp(toAddress, toName, subject string, contents io.Reader) error {
	host, ok := os.LookupEnv("MAIL_HOST")
	if !ok {
		return errors.New("environment variable MAIL_HOST not set")
	}
	portStr, ok := os.LookupEnv("MAIL_PORT")
	if !ok {
		return errors.New("environment variable MAIL_PORT not set")
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return fmt.Errorf("cannot convert mailserver port number %q to int", portStr)
	}
	username, ok := os.LookupEnv("MAIL_USERNAME")
	if !ok {
		return errors.New("environment variable MAIL_USERNAME not set")
	}
	password, ok := os.LookupEnv("MAIL_PASSWORD")
	if !ok {
		return errors.New("environment variable MAIL_PASSWORD not set")
	}
	fromAddress, ok := os.LookupEnv("MAIL_FROM_ADDRESS")
	if !ok {
		return errors.New("environment variable MAIL_FROM_ADDRESS not set")
	}
	fromName, ok := os.LookupEnv("MAIL_FROM_NAME")
	if !ok {
		return errors.New("environment variable MAIL_FROM_NAME not set")
	}
	body, err := ioutil.ReadAll(contents)
	if err != nil {
		return fmt.Errorf("could not read contents of contents reader: %w", err)
	}

	m := gomail.NewMessage()
	m.SetAddressHeader("From", fromAddress, fromName)
	m.SetAddressHeader("To", toAddress, toName)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", string(body))

	d := gomail.NewDialer(host, port, username, password)
	return d.DialAndSend(m)
}

func sendSignupEmail(u *ent.User) error {
	tmpl, err := template.New("signup.html").ParseFiles("resources/emails/signup.html")
	if err != nil {
		return fmt.Errorf("could not parse signup email template: %w", err)
	}
	confirmUrl, err := absoluteUrl(fmt.Sprintf("confirm_email?email=%v&token=%v", url.QueryEscape(u.Email), u.VerifyToken))
	if err != nil {
		return fmt.Errorf("could not get confirmation url: %w", err)
	}
	var contents bytes.Buffer
	err = tmpl.Execute(&contents, struct {
		Username string
		Url      string
	}{
		Username: u.Username,
		Url:      confirmUrl,
	})
	if err != nil {
		return fmt.Errorf("could not execute signup email template: %w", err)
	}

	return sendEmail(
		u.Email,
		u.Username,
		"Welcome to Empire and Foundation",
		&contents,
	)
}
