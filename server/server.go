package server

import (
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"strconv"
)

func Router() {

	router := httprouter.New()
	router.GET("/", Index)
	//	router.GET("/api/*path", ApiGET)
	//	router.PUT("/api/*path", ApiPUT)
	//	router.POST("/api/*path", ApiPOST)
	//	router.DELETE("/api/*path", ApiDELETE)
	log.Debug("Router initialized")
	router.ServeFiles("/static/*filepath", http.Dir("static/"))

	CertificateFile := viper.GetString("certificate-file")
	KeyFile := viper.GetString("key-file")
	Port := viper.GetInt("port")

	if _, err := os.Stat(CertificateFile); err == nil {
		if _, err := os.Stat(KeyFile); os.IsNotExist(err) {
			log.Error("Found " + CertificateFile + " but missing " + KeyFile)
			os.Exit(10)
		}
		log.Fatal(http.ListenAndServeTLS(":"+strconv.Itoa(Port), CertificateFile, KeyFile, router))

	} else {
		log.Fatal(CertificateFile + " not found")
		os.Exit(10)
	}
}

func RequestLogger(R *http.Request, Message string) {
	log.WithFields(log.Fields{
		"method": R.Method,
		"url":    R.URL.String(),
		"remote": R.RemoteAddr,
		"proto":  R.Proto,
	}).Info(Message)
}
