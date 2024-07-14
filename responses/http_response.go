package responses

import (
	"encoding/json"
	"net/http"

	models "GitHub.com/weatherAPI/models/json"
)

// SSends a JSON response with success status and data
func SendJSONSuccess(w http.ResponseWriter, status int, data interface{}) {
	sendJSONResponse(w, status, false, "Success", data)
}

// Sends a JSON response with error status and message
func SendJSONError(w http.ResponseWriter, status int, message string) {
	sendJSONResponse(w, status, true, message, []models.City{})
}

// Sends a JSON response with the provided parameters
func sendJSONResponse(w http.ResponseWriter, status int, isError bool, msg string, responseData interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := models.Response{
		Status: status,
		Error:  isError,
		Msg:    msg,
		Data:   responseData,
	}

	json.NewEncoder(w).Encode(response)
}
