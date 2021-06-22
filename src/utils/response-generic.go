package utils

import (
	"encoding/json"
	"net/http"
)

type responseGenric struct {
	Message string `json:"message"`
	Code int `json:"code"`
}

var response = responseGenric{
	Message: "Operation successfully",
	Code: 200,
}

func ResponseGenericOK(w http.ResponseWriter){
	json.NewEncoder(w).Encode(response)
}