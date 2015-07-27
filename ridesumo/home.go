package ridesumo

import (
	"appengine"
	"github.com/gorilla/context"
	"github.com/strava/go.strava"
	"log"
	"middleware"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	user := context.Get(r, middleware.ContextUserKey).(*User)
	if user != nil && user.StravaAccessToken != "" {
		log.Printf("redirecting logged '%s' to dashboard", user.Email)
		http.Redirect(w, r, "/u/dashboard", http.StatusFound)
		return
	}

	log.Println("offering login with Strava")
	renderer.HTML(w, http.StatusOK, "home", map[string]interface{}{
		"authURL": stravaOAuth.AuthorizationURL("login", strava.Permissions.Public, appengine.IsDevAppServer()),
	})
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	user := context.Get(r, middleware.ContextUserKey).(*User)

	comps, _ := GetAllComps()
	renderer.HTML(w, http.StatusOK, "dashboard", map[string]interface{}{
		"user":  user,
		"comps": comps,
	})
}

func GetComps(w http.ResponseWriter, r *http.Request) {
	//
}
func GetComp(w http.ResponseWriter, r *http.Request) {
}
func NewComp(w http.ResponseWriter, r *http.Request) {
}
func CreateComp(w http.ResponseWriter, r *http.Request) {
	// use the Accept-Encoding header to control whether to respond with JSON or HTML
}
func EditComp(w http.ResponseWriter, r *http.Request) {
	// use the Accept-Encoding header to control whether to respond with JSON or HTML
}
