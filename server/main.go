// server is the REST API and webserver component of the ppapt-backend
//
// This file, main.go, contains the main router and functions used throughout
// the package.
package server

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/ppapt/ppapt-backend/ppapt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"strconv"
)

// The main ppapt object handles all application logic
var p *ppapt.Ppapt

// Router opens a web server and routes calls to the static web server or the
// various API functions
func Router(Ppapt *ppapt.Ppapt) {
	logger := log.WithFields(log.Fields{
		"context": "server",
		"func":    "Router",
	})
	p = Ppapt
	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/api/login", ApiLogin)
	router.POST("/api/logout", ApiLogout)
	//	router.PUT("/api/*path", ApiPUT)
	//	router.POST("/api/*path", ApiPOST)
	//	router.DELETE("/api/*path", ApiDELETE)
	logger.Debug("Router initialized")
	router.ServeFiles("/static/*filepath", http.Dir("static/"))

	CertificateFile := viper.GetString("certificate-file")
	KeyFile := viper.GetString("key-file")
	Port := viper.GetInt("port")

	if _, err := os.Stat(CertificateFile); err == nil {
		if _, err := os.Stat(KeyFile); os.IsNotExist(err) {
			logger.Fatal("Found " + CertificateFile + " but missing " + KeyFile)
			os.Exit(10)
		}
		logger.Fatal(http.ListenAndServeTLS(":"+strconv.Itoa(Port), CertificateFile, KeyFile, router))

	} else {
		logger.Fatal(CertificateFile + " not found")
		os.Exit(10)
	}
}

// RequestLogger logs a http request
// R is the http request
// Message the log message to be output at Info log level
func RequestLogger(R *http.Request, Message string) {
	log.WithFields(log.Fields{
		"method": R.Method,
		"url":    R.URL.String(),
		"remote": R.RemoteAddr,
		"proto":  R.Proto,
	}).Info(Message)
}

// sendJSON sends a JSON response out of the provided data
func sendJSON(w http.ResponseWriter, data interface{}, status int) {
	j, err := json.Marshal(data)
	if err != nil {
		log.WithFields(log.Fields{
			"func":   sendJSON,
			"action": "json.Marshal",
		}).Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(j))
}
