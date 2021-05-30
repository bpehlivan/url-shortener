package api

import (
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"src/models"

	"github.com/gorilla/mux"
)

func dbWrapper(handlerFunc func(*gorm.DB, http.ResponseWriter, *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	dbWrappedFunc := func (rw http.ResponseWriter, r *http.Request) {
		gormDb := models.GetConnection()
		handlerFunc(gormDb, rw, r)
		models.CloseConnection(gormDb)
	}
	return dbWrappedFunc
}



func SetUrlHandlers(router *mux.Router) {
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received.")
	})
	router.HandleFunc("/urls/{endpointKeyword}", dbWrapper(retrieveUrl)).Methods("GET")
	router.HandleFunc("/urls/", dbWrapper(createUrl)).Methods("POST")
	router.Use(loggingMiddleware)
}
