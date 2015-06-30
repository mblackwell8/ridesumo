package ridesumo

import (
	// "fmt"
	// "github.com/codegangsta/negroni"
	// "github.com/gorilla/context"
	// "github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
)

var renderer = render.New(render.Options{})

func Test(w http.ResponseWriter, r *http.Request) {
	renderer.JSON(w, http.StatusOK, map[string]string{
		"test": "OK",
	})
}

func GetLogin(w http.ResponseWriter, r *http.Request) {
}
func Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Please login"))
}
func ConnectWithStrava(w http.ResponseWriter, r *http.Request) {
	// redirect to Strava...
	// https://www.strava.com/oauth/authorize?client_id=9&response_type=code&redirect_uri=http://localhost/token_exchange.php&approval_prompt=force
}
