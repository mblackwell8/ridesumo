package ridesumo

import (
	// "fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	// "github.com/gorilla/sessions"
	"net/http"
	"sessionauth"
)

func init() {
	app := negroni.Classic()

	// store := sessions.NewCookieStore([]byte("something-very-secret"))
	// app.Use(sessionauth.SessionUser(store, "rs_sess", GenerateAnonymousUser))

	app.UseHandler(makeRoutes())

	http.Handle("/", context.ClearHandler(app))
}

func makeRoutes() http.Handler {
	router := mux.NewRouter()
	// router.HandleFunc("/", TempRedirect)
	router.HandleFunc("/test", Test).Methods("GET")
	router.HandleFunc("/login", Login).Methods("GET")

	userRoutes := mux.NewRouter()
	userRoutes.HandleFunc("/test", Test).Methods("GET")

	// Create a new negroni for the user middleware
	router.Handle("/user", negroni.New(
		negroni.HandlerFunc(sessionauth.LoginRequired),
		negroni.Wrap(userRoutes),
	))

	return router
}
