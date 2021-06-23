package routes

import (
	"api-rest/src/errors"
	"api-rest/src/middleware"
	"api-rest/src/models/dtos"
	"api-rest/src/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type LoginDto = dtos.LoginDto

func AuthRoute(router *mux.Router){
	router.HandleFunc("/login",middleware.SetHeaderJSON(login)).Methods("POST")
}

func login(w http.ResponseWriter, r *http.Request){
	var loginDto LoginDto
	if err := utils.Deserialize(r.Body,&loginDto); err != nil {
		errors.HandleError(w,*err)
		return
	}
	/*var users, e = services.GetUsers()
	if(e != nil){
		errors.HandleError(w,*e)
		return
	}
	json.NewEncoder(w).Encode(users)*/
}

