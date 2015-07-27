package middleware

import (
	"appengine"
	"log"
	"net/http"
)

var AppEngineContext appengine.Context

func CreateAppEngineContext(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("setting appengine context")
	AppEngineContext = appengine.NewContext(r)

	next(rw, r)
}
