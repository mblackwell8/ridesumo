package ridesumo

import (
	// "fmt"
	"appengine"
	"appengine/urlfetch"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/strava/go.strava"
	"github.com/unrolled/render"
	"log"
	"middleware"
	"net/http"
)

var renderer = render.New(render.Options{
	IndentJSON: true,
})

// default sessions last for one month
var store = sessions.NewCookieStore([]byte("something-very-secret"))
var sessionName = "rs_sess"

func init() {
	app := negroni.New(
		negroni.NewRecovery(), 
		negroni.NewLogger(), 
		negroni.NewStatic(http.Dir("static")))

	log.Println("setting AE middleware")
	// must be done before session because it needs the appengine context
	app.Use(negroni.HandlerFunc(middleware.CreateAppEngineContext))

	log.Println("setting session middleware")
	app.Use(middleware.SessionUser(store, sessionName, GenerateAnonymousUser))

	initStravaOAuth()
	app.UseHandler(makeRoutes())

	http.Handle("/", context.ClearHandler(app))
}

func makeRoutes() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/test", Test).Methods("GET")
	router.HandleFunc("/register", Register).Methods("POST")

	path, _ := stravaOAuth.CallbackPath()
	router.HandleFunc(path, stravaOAuth.HandlerFunc(oAuthSuccess, oAuthFailure))

	userRoutes := mux.NewRouter()
	userRoutes.HandleFunc("/u/dashboard", Dashboard).Methods("GET")
	userRoutes.HandleFunc("/u/comps", GetComps).Methods("GET")
	userRoutes.HandleFunc("/u/comps/{id}", GetComp).Methods("GET")
	userRoutes.HandleFunc("/u/comps", CreateComp).Methods("POST")
	userRoutes.HandleFunc("/u/comps/{id}", EditComp).Methods("POST")

	router.PathPrefix("/u").Handler(negroni.New(
		negroni.HandlerFunc(middleware.LoginRequired),
		negroni.Wrap(userRoutes),
	))

	// router.HandleFunc("/login", GetLogin).Methods("GET")
	// router.HandleFunc("/login", Login).Methods("POST")

	// userRoutes := router.Path("/u").Subrouter()
	// router.HandleFunc("/u/test", Test).Methods("GET")
	// userRoutes.HandleFunc("/dashboard", Test).Methods("GET")

	// Create a new negroni for the user middleware
	// router.Handle("/u", negroni.New(
	// 	negroni.HandlerFunc(sessionauth.LoginRequired),
	// 	negroni.Wrap(userRoutes),
	// ))

	return router
}

var stravaOAuth *strava.OAuthAuthenticator

func initStravaOAuth() {
	strava.ClientId = 6897
	strava.ClientSecret = "b63e66cfc9f1280d0ad5fa3d7eeeb40e10628cba"
	stravaOAuth = &strava.OAuthAuthenticator{
		CallbackURL: "http://localhost:8080/exchange_token",
		RequestClientGenerator: func(r *http.Request) *http.Client {
			return urlfetch.Client(appengine.NewContext(r))
		},
	}
}

func Test(w http.ResponseWriter, r *http.Request) {
	renderer.HTML(w, http.StatusOK, "default", nil)
	// renderDefaultHTML(w, "default", nil)
}
