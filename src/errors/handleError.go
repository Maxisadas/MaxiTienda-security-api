package errors

import (
	"encoding/json"
	"net/http"
)

type HTTPErrors struct {
	Message string `json:"message"`
	Code int `json:"code"`
}

func HandleError(w http.ResponseWriter, e HTTPErrors){
	w.WriteHeader(e.Code)
	json.NewEncoder(w).Encode(e)
}