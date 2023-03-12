package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func BParses(r *http.Request) []byte {
	body, _ := ioutil.ReadAll(r.Body)
	return body
}

func ToJSON(w http.ResponseWriter, data interface{}, statusCode int)  {
	w.Header().Set("Content-Type", "application/json; charset=UTF8")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	CheckErr(err)
}