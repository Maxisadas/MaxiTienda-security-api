package routes

import (
	"api-rest/src/errors"
	"api-rest/src/middleware"
	"api-rest/src/models/dtos"
	"api-rest/src/services"
	"api-rest/src/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CreateUserDto = dtos.CreateUserDto;

func UserRoute(router *mux.Router){
	router.HandleFunc("/users",middleware.SetHeaderJSON(getUsers)).Methods("GET")
	router.HandleFunc("/users",middleware.SetHeaderJSON(createUser)).Methods("POST")
}

func getUsers(w http.ResponseWriter, r *http.Request){
	var users, e = services.GetUsers()
	if(e != nil){
		errors.HandleError(w,*e)
		return
	}
	json.NewEncoder(w).Encode(users)	
}

func createUser(w http.ResponseWriter, r *http.Request){
	var user CreateUserDto
	if err := utils.Deserialize(r.Body,&user); err != nil {
		errors.HandleError(w,*err)
		return
	}
	var e = services.CreateUser(user)
	if(e != nil){
		errors.HandleError(w,*e)
		return
	}
	utils.ResponseGenericOK(w);
}
