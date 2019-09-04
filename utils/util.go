
package utils

import (
	"encoding/json"
	"net/http"
	"log"
)

func Message(status bool, message string) (map[string]interface{}) {
	return map[string]interface{} {"status" : status, "message" : message}
}

func Respond(w http.ResponseWriter, data interface{})  {
	log.Println("Data: ",data)
	 w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}