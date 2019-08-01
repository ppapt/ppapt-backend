// server is the REST API and webserver component of the ppapt-backend
//
// This file, api.go, implements the REST API
package server

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"github.com/ppapt/ppapt-backend/data"
)

// ApiLogin handles a User login and returns a token to the caller. If the
// user can't be retrieved or the password doesn't match, A 401 is returned.
func ApiLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var in data.LoginData

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
	user, err := DB.GetUser(in.EMail)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !user.PasswordIsValid(in.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	t := Session.AddSession(in.EMail)
	tr := data.TokenResponse{
		Token: t,
	}
	sendJSON(w, tr, http.StatusOK)
}
