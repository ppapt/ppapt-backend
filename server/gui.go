// server is the REST API and webserver component of the ppapt-backend
//
// This file, gui.go, implements gui related functions
package server

import (
	"github.com/julienschmidt/httprouter"
	//log "github.com/sirupsen/logrus"
	"net/http"
)

// Index redirects the caller to the static index page
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	RequestLogger(r, "Index")
	http.Redirect(w, r, "/static/index.html", http.StatusPermanentRedirect)
}
