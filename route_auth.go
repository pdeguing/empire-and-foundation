package main

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/pdeguing/empire-and-foundation/data"
	"github.com/pdeguing/empire-and-foundation/ent"
	"github.com/pdeguing/empire-and-foundation/ent/planet"
	"github.com/pdeguing/empire-and-foundation/ent/user"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"math/rand"
	"net/http"
)

// GET /signup
// Show the signup page
func serveSignup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, r, "signup", nil, "layout", "public.navbar", "flash", "signup")
	forgetForm(r)
	forgetFormErrors(r)
}

type signupAccountRequest struct {
	Email          string `json:"email" name:"email" validate:"required,email"`
	Username       string `json:"username" name:"username" validate:"required,min=2,max=30,unique_username"`
	Password       string `json:"password" name:"password" validate:"required,min=8"`
	PasswordRepeat string `json:"password_confirm" name:"confirm password" validate:"required,eqfield=Password"`
}

// POST /signup
// Create the user account
func serveSignupAccount(w http.ResponseWriter, r *http.Request) {
	forgetFormErrors(r)
	var input signupAccountRequest
	if err := r.ParseForm(); err != nil {
		serveError(w, r, newInternalServerError(fmt.Errorf("unable to parse form: %v", err)))
		return
	}
	if err := decoder.Decode(&input, r.PostForm); err != nil {
		serveError(w, r, newInternalServerError(fmt.Errorf("unable to decode form: %v", err)))
		return
	}
	if err := validate.StructCtx(r.Context(), input); err != nil {
		var valErrs validator.ValidationErrors
		if errors.As(err, &valErrs) {
			storeFormErrors(r, valErrs)
			rememberForm(r)
			http.Redirect(w, r, "/signup", 303)
			return
		}
		serveError(w, r, newInternalServerError(fmt.Errorf("could not validate the form: %v", err)))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	if err != nil {
		serveError(w, r, newInternalServerError(fmt.Errorf("unable to encrypt password: %v", err)))
		return
	}

	verifyToken, err := generateRandomString(20)
	if err != nil {
		serveError(w, r, newInternalServerError(fmt.Errorf("unable to generate verify token: %v", err)))
		return
	}

	err = data.WithTx(r.Context(), data.Client, func(tx *ent.Tx) error {
		// Don't create an account for the user if an account with this email already exists.
		// For privacy reasons we can't respond with a form error because that would disclose
		// who registered to the game and who didn't. Instead we can e-mail the user a link
		// to reset their password.
		// TODO: first implement a password reset. Then send a modified password reset email
		//       when the user tries to register with an email that is already used. But
		//       change the message a bit so that it is clear why they received a password
		//       reset instead of an activation email and that they can ignore the email
		//       if it wasn't them.
		n, err := tx.User.
			Query().
			Where(user.EmailEQ(r.PostFormValue("email"))).
			Count(r.Context())
		if n > 0 {
			return nil
		}

		u, err := tx.User.
			Create().
			SetUsername(r.PostFormValue("name")).
			SetEmail(r.PostFormValue("email")).
			SetPassword(string(hashedPassword)).
			SetVerifyToken(verifyToken).
			Save(r.Context())
		if err != nil {
			return err
		}

		c, err := tx.Planet.Query().
			Where(
				planet.And(
					planet.PlanetTypeEQ(planet.PlanetTypeHabitable),
					planet.Not(planet.HasOwner()),
				),
			).
			Count(r.Context())

		if err != nil {
			return err
		}

		ra := rand.New(rand.NewSource(time.Now().UnixNano()))
		rn := ra.Intn(c)

		p, err := tx.Planet.Query().
			Where(
				planet.And(
					planet.PlanetTypeEQ(planet.PlanetTypeHabitable),
					planet.Not(planet.HasOwner()),
				),
			).
			Offset(rn).
			First(r.Context())

		_, err = p.Update().
			SetOwner(u).
			Save(r.Context())

		if err := sendSignupEmail(u); err != nil {
			return err
		}

		return err
	})
	if err != nil {
		serveError(w, r, newInternalServerError(fmt.Errorf("unable to create a user account: %v", err)))
		return
	}

	flash(r, flashSuccess, "Your account has been created. You can log in now.")
	http.Redirect(w, r, "/", 303)
}

func sendSignupEmail(u *ent.User) error {
	tmpl, err := template.New("signup.html").ParseFiles("resources/emails/signup.html")
	if err != nil {
		return fmt.Errorf("could not parse signup email template: %w", err)
	}
	var contents bytes.Buffer
	err = tmpl.Execute(&contents, struct {
		Username string
		Url      string
	}{
		Username: u.Username,
		Url:      fmt.Sprintf("https://example.com/?token=%v", u.VerifyToken),
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

// POST /authenticate
// Authenticate the user given the email and password
func serveAuthenticate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		serveError(w, r, newInternalServerError(fmt.Errorf("unable to parse form: %v", err)))
		return
	}
	u, err := data.Client.User.
		Query().
		Where(user.Email(r.PostFormValue("email"))).
		Only(r.Context())
	var nferr *ent.NotFoundError
	if errors.As(err, &nferr) {
		flash(r, flashDanger, "The username or password you have entered is invalid.")
		rememberForm(r)
		http.Redirect(w, r, "/", 302)
		return
	}
	if err != nil {
		serveError(w, r, newInternalServerError(fmt.Errorf("unable to check if user exists in database: %v", err)))
		return
	}
	ok, err := data.CheckPassword(u.Password, r.PostFormValue("password"))
	if err != nil {
		serveError(w, r, newInternalServerError(fmt.Errorf("unable to validate the users password: %v", err)))
		return
	}
	if !ok {
		flash(r, flashDanger, "The username or password you have entered is invalid.")
		rememberForm(r)
		http.Redirect(w, r, "/", 302)
		return
	}

	authenticate(r, u)
	http.Redirect(w, r, "/dashboard", 302)
}

// GET /logout
// Logs the user out
func serveLogout(w http.ResponseWriter, r *http.Request) {
	logout(r)
	http.Redirect(w, r, "/", 302)
}
