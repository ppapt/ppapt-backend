// server is the REST API and webserver component of the ppapt-backend
//
// This file, api.go, implements the REST API
package server

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"github.com/ppapt/ppapt-backend/ppapt"
)

// ApiLogin handles a User login and returns a token to the caller. If the
// user can't be retrieved or the password doesn't match, A 401 is returned.
func ApiLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var in ppapt.LoginData

	RequestLogger(r, "api/login")
	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(response, &in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, t, err := p.Login(in.EMail,in.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	tr := ppapt.TokenResponse{
		Token: t,
		User: user,
	}
	sendJSON(w, tr, http.StatusOK)
}
