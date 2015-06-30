package sessionauth

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	// "log"
	"github.com/codegangsta/negroni"
	"net/http"
)

// These are the default configuration values for this package. They
// can be set at anytime, probably during the initial setup of Martini.
var (
	// RedirectUrl should be the relative URL for your login route
	RedirectUrl string = "/login"

	// RedirectParam is the query string parameter that will be set
	// with the page the user was trying to visit before they were
	// intercepted.
	RedirectParam string = "next"

	// SessionKey is the key containing the unique ID in your session
	SessionKey string = "AUTHUNIQUEID"

	UserKey string = "context-user"
)

// User defines all the functions necessary to work with the user's authentication.
// The caller should implement these functions for whatever system of authentication
// they choose to use
type User interface {
	// // Return whether this user is logged in or not
	IsAuthenticated() bool

	// Set any flags or extra data that should be available
	Login()

	// Clear any sensitive data out of the user
	Logout()

	// Return the unique identifier of this user object
	UniqueId() interface{}

	// Populate this user object with values
	GetById(id interface{}) error
}

// SessionUser will try to read a unique user ID out of the session. Then it tries
// to populate an anonymous user object from the database based on that ID. If this
// is successful, the valid user is mapped into the context. Otherwise the anonymous
// user is mapped into the contact.
// The newUser() function should provide a valid 0value structure for the caller's
// user type.
func SessionUser(store sessions.Store, sessionName string, newUser func() User) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		session, _ := store.Get(r, sessionName)
		userId := session.Values[SessionKey]
		user := newUser()

		if userId != nil {
			err := user.GetById(userId)
			if err != nil {
				// l.Printf("Login Error: %v\n", err)
			} else {
				user.Login()
			}
		}

		context.Set(r, UserKey, user)
	}
}

// called by the Login process

// AuthenticateSession will mark the session and user object as authenticated. Then
// the Login() user function will be called. This function should be called after
// you have validated a user.
func AuthenticateSession(s sessions.Session, user User) error {
	user.Login()
	return UpdateUser(s, user)
}

// Logout will clear out the session and call the Logout() user function.
func Logout(s sessions.Session, user User) {
	user.Logout()
	delete(s.Values, SessionKey)
}

// LoginRequired verifies that the current user is authenticated. Any routes that
// require a login should have this handler placed in the flow. If the user is not
// authenticated, they will be redirected to /login with the "next" get parameter
// set to the attempted URL.
// func LoginRequired(r render.Render, user User, req *http.Request) {
// 	if user.IsAuthenticated() == false {
// 		path := fmt.Sprintf("%s?%s=%s", RedirectUrl, RedirectParam, req.URL.Path)
// 		r.Redirect(path, 302)
// 	}
// }

// If a middleware hasn't already written to the ResponseWriter, it should call the next http.HandlerFunc in the chain to yield to the next middleware handler.
func LoginRequired(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// TODO: check the type
	user := context.Get(r, UserKey).(User)
	if user.IsAuthenticated() == false {
		path := fmt.Sprintf("%s?%s=%s", RedirectUrl, RedirectParam, r.URL.Path)
		http.Redirect(w, r, path, http.StatusFound)
	} else {
		next(w, r)
	}
}

// UpdateUser updates the User object stored in the session. This is useful incase a change
// is made to the user model that needs to persist across requests.
func UpdateUser(s sessions.Session, user User) error {
	s.Values[SessionKey] = user.UniqueId()
	return nil
}
