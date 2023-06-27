package util

import (
	"encoding/json"
	"net/http"
)

func WriteResponseJson(writer http.ResponseWriter, response interface{}, statusCode ...int) {
	writer.Header().Set("Content-Type", "application/json")
	if len(statusCode) > 0 {
		writer.WriteHeader(statusCode[0])
	} else {
		writer.WriteHeader(http.StatusOK)
	}
	if response != nil {
		json.NewEncoder(writer).Encode(response)
	}
}
