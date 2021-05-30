package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"src/models"
)


func retrieveUrl(db *gorm.DB, rw http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	endpointKeyword := vars["endpointKeyword"]
	fmt.Printf("endpoint keyword: %s \n", endpointKeyword)
	endpoint := models.EndPoint{}
	result := db.Where(&models.EndPoint{KeyWord: endpointKeyword}).First(&endpoint)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound){
			responseJsonError(rw, http.StatusNotFound, "Record not found")
		} else {
			responseJsonError(rw, http.StatusInternalServerError, "Internal Server Error")
		}
	}

	http.Redirect(rw, request, endpoint.Url, http.StatusSeeOther)
}

type createUrlInput struct {
	Url string
}

func createUrl(db *gorm.DB, rw http.ResponseWriter, request *http.Request) {
	// pre-validations
	input := createUrlInput{}
	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&input); err != nil {
		responseJsonError(rw, http.StatusBadRequest, err.Error())
		return
	}

	if input.Url == "" {
		responseJsonError(rw, http.StatusBadRequest, "\"Url\" field is required.")
	}
	if err := request.Body.Close(); err != nil {
		errMsg := fmt.Sprintf("Error at file closing \n %s", err.Error())
		panic(errMsg)
	}

	// action
	endpoint := models.EndPoint{Url: input.Url}
	result := db.Create(&endpoint)
	if result.Error != nil {
		responseJsonError(rw, http.StatusInternalServerError, "Internal Server Error")
	}

	jsonResponse(rw, http.StatusCreated, endpoint)
}
