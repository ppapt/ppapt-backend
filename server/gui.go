package server

import (
	"github.com/julienschmidt/httprouter"
	//log "github.com/sirupsen/logrus"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	RequestLogger(r, "Index")
	http.Redirect(w, r, "/static/index.html", http.StatusPermanentRedirect)
}
