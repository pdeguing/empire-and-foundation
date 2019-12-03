package main

import (
	"context"
	"encoding/gob"
	"net/http"
	"net/url"
	"time"

	"github.com/alexedwards/scs"
	"github.com/alexedwards/scs/postgresstore"
	"github.com/pdeguing/empire-and-foundation/data"
)

func init() {
	// Register all types that can be stored in the session
	// so they can be serialized.
	gob.Register(&flashMessage{})
	gob.Register(&url.Values{}) // Form fields
}

var sessionManager = func() *scs.SessionManager {
	mngr := scs.New()
	mngr.Lifetime = 24 * time.Hour
	mngr.Store = postgresstore.New(data.Db)
	return mngr
}()

// key is an unexported type for keys defined in this package.
// This prevents collisions with keys defined in other packages.
type key int

// userKey is the key for user.User values in Contexts. It is
// unexported; clients use user.NewContext and user.FromContext
// instead of using this key directly.
var userKey key

// Session keys
const (
	flashMessageKey = "flash_message"
	formFieldsKey   = "form_fields"
	userIDKey       = "user_id"
)

// authMiddleware checks that the user is authenticated before
// proceeding with the request. If the user is not authenticated,
// they will be redirected to the login page. authMiddleware also
// adds the user object to the request's context.
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if !sessionManager.Exists(ctx, userIDKey) {
			http.Redirect(w, r, "/login", 302)
			return
		}

		userID := sessionManager.GetInt(ctx, userIDKey)
		user, err := data.UserByID(userID)
		if err != nil {
			internalServerError(w, r, err, "Could not load logged in user from database")
			return
		}

		ctx = context.WithValue(ctx, userKey, &user)
		sr := r.WithContext(ctx)
		next.ServeHTTP(w, sr)
	})
}

// renewSessionToken changes the token used to store the session.
// This should be done after any privilege level change like logging
// in. This will prevent session fixation attacks.
func renewSessionToken(r *http.Request) error {
	return sessionManager.RenewToken(r.Context())
}

type flashType byte

const (
	flashInfo flashType = iota
	flashSuccess
	flashWarning
	flashDanger
)

type flashMessage struct {
	Message string
	Type    flashType
}

func flash(r *http.Request, typ flashType, message string) {
	sessionManager.Put(r.Context(), flashMessageKey, &flashMessage{
		Message: message,
		Type:    typ,
	})
}

func hasFlash(r *http.Request) bool {
	return sessionManager.Exists(r.Context(), flashMessageKey)
}

func getFlash(r *http.Request) *flashMessage {
	val := sessionManager.Pop(r.Context(), flashMessageKey)
	if val == nil {
		return nil
	}
	f, ok := val.(*flashMessage)
	if !ok {
		panic("Could not cast the flash message to the flashMessage type")
	}
	return f
}

// rememberForm stores the values of the form fields in the session
// so that they can be pre-filled when the form is rendered again
func rememberForm(r *http.Request) {
	sessionManager.Put(r.Context(), formFieldsKey, &r.PostForm)
}

// forgetForm removes the previously stored form field values from
// the session. This should always be done after a page is rendered.
func forgetForm(r *http.Request) {
	sessionManager.Remove(r.Context(), formFieldsKey)
}

// oldFormValue retrieves the old form field value from the session
// or an empty string if it isn't available
func oldFormValue(r *http.Request, field string) string {
	val := sessionManager.Get(r.Context(), formFieldsKey)
	if val == nil {
		return ""
	}
	f, ok := val.(*url.Values)
	if !ok {
		panic("Could not cast the form fields to the url.Values type")
	}
	return f.Get(field)
}

// authenticate sets user as the currently logged in user.
// The change will take effect on the next request.
func authenticate(r *http.Request, user *data.User) {
	renewSessionToken(r)
	sessionManager.Put(r.Context(), userIDKey, user.Id)
}

// logout logs the user out.
// The change will take effect on the next request.
func logout(r *http.Request) {
	renewSessionToken(r)
	sessionManager.Remove(r.Context(), userIDKey)
}

// isAuthenticated checks if the user is logged in
func isAuthenticated(r *http.Request) bool {
	return r.Context().Value(userKey) != nil
}

// user returns the currently logged in user, or nil if the user
// isn't logged in
func user(r *http.Request) *data.User {
	val := r.Context().Value(userKey)
	if val == nil {
		return nil
	}
	user, ok := val.(*data.User)
	if !ok {
		panic("Unable to cast the user object stored in the context to a *data.User")
	}
	return user
}
