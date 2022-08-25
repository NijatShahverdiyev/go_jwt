package util

import (
	"encoding/json"
	"log"
	"net/http"
)

func SetHeader(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
}
func ErrorHandle(w http.ResponseWriter, data ErrorResponse) {
	err := json.NewEncoder(w).Encode(data)
	Notify(err, data.Message)
}

func Notify(err error, message string) {
	if err != nil {
		log.Fatalf("%v %v", err, message)
	}
	return
}
