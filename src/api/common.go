package api

import (
	"encoding/json"
	"net/http"
)

// jsonResponse makes the response with payload as json format
func jsonResponse(rw http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write([]byte(response))
}

// responseJsonError makes the error response with payload as json format
func responseJsonError(w http.ResponseWriter, code int, message string) {
	jsonResponse(w, code, map[string]string{"error": message})
}
