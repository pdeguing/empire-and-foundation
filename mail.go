package main

import (
	"errors"
	"fmt"
	"gopkg.in/gomail.v2"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func sendEmail(toAddress, toName, subject string, contents io.Reader) error {
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
