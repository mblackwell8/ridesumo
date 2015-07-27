package ridesumo

import (
	"fmt"
	// "log"
	// "github.com/codegangsta/negroni"
	// "github.com/gorilla/context"
	// "github.com/gorilla/mux"
	// "encoding/json"
	// "github.com/gorilla/sessions"
	"github.com/strava/go.strava"
	"middleware"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
}

// func GetLogin(w http.ResponseWriter, r *http.Request) {
// 	renderer.HTML(w, http.StatusOK, "login", nil)
// }
// func Login(w http.ResponseWriter, r *http.Request) {
// }

func oAuthSuccess(auth *strava.AuthorizationResponse, w http.ResponseWriter, r *http.Request) {
	// create the user
	// HACK: are these mandatory from Strava? probably
	user := NewUser(auth.Athlete.Email, auth.Athlete.FirstName, auth.Athlete.LastName)
	user.City = auth.Athlete.City
	user.State = auth.Athlete.State
	user.Country = auth.Athlete.Country
	user.Gender = string(auth.Athlete.Gender)
	user.StravaAccessToken = auth.AccessToken
	// user.StravaUserRecord = auth.Athlete

	// save the user
	err := user.Put()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// authenticate the session (ie. store the user's unique ID)
	session, _ := store.Get(r, sessionName)
	middleware.AuthenticateSession(session, user)
	session.Save(r, w)

	// fmt.Fprintf(w, "SUCCESS:\nAt this point you can use this information to create a new user or link the account to one of your existing users\n")
	// fmt.Fprintf(w, "State: %s\n\n", auth.State)
	// fmt.Fprintf(w, "Access Token: %s\n\n", auth.AccessToken)

	// fmt.Fprintf(w, "The Authenticated Athlete (you):\n")
	// content, _ := json.MarshalIndent(auth.Athlete, "", " ")
	// fmt.Fprint(w, string(content))

	// redirect to dashboard
	// log.Printf("setting user in local context: %+v", user)
	// context.Set(r, sessionauth.UserKey, user)
	http.Redirect(w, r, "/u/dashboard", http.StatusFound)
}

func oAuthFailure(err error, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Authorization Failure:\n")

	// some standard error checking
	if err == strava.OAuthAuthorizationDeniedErr {
		fmt.Fprint(w, "The user clicked the 'Do not Authorize' button on the previous page.\n")
		fmt.Fprint(w, "This is the main error your application should handle.")
	} else if err == strava.OAuthInvalidCredentialsErr {
		fmt.Fprint(w, "You provided an incorrect client_id or client_secret.\nDid you remember to set them at the begininng of this file?")
	} else if err == strava.OAuthInvalidCodeErr {
		fmt.Fprint(w, "The temporary token was not recognized, this shouldn't happen normally")
	} else if err == strava.OAuthServerErr {
		fmt.Fprint(w, "There was some sort of server error, try again to see if the problem continues")
	} else {
		fmt.Fprint(w, err)
	}
}
