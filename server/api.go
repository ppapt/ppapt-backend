// server is the REST API and webserver component of the ppapt-backend
//
// This file, api.go, implements the REST API
package server

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/ppapt/ppapt-backend/ppapt"
	"github.com/ppapt/ppapt-backend/common"
	"io/ioutil"
	"net/http"
)

// ApiLogin handles a User login and returns a token to the caller. If the
// user can't be retrieved or the password doesn't match, A 401 is returned.
func ApiLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var in common.LoginData

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
	user, t, err := p.Login(in.EMail, in.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	tr := ppapt.TokenResponse{
		Token: t,
		User:  user,
	}
	emailcookie:=http.Cookie{Name: "email",Value: user.EMail}
    http.SetCookie(w, &emailcookie)
	tokencookie:=http.Cookie{Name: "token",Value: t}
    http.SetCookie(w, &tokencookie)
	sendJSON(w, tr, http.StatusOK)
}

// ApiLogout removes the users session from the session list
func ApiLogout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var EMail string
	var Token string
	for _, cookie := range r.Cookies() {
		if cookie.Name == "email" {
			EMail=cookie.Value
		}
		if cookie.Name == "token" {
			Token=cookie.Value
		}
	}
	_,err:=p.GetSession(EMail,Token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p.DeleteSession(EMail)
	m:=common.Message{
		Message: "Logged out successfully",
	}
	sendJSON(w,m,http.StatusOK)
}
