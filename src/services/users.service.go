package services

import (
	db "api-rest/src/databases"
	"api-rest/src/errors"
	"api-rest/src/models"
	"api-rest/src/models/dtos"

	"golang.org/x/crypto/bcrypt"
)

type User = models.User
type HTTPError = errors.HTTPErrors


func GetUsers() ([]User,*HTTPError){
	var users []User
	result := db.DBConn.Find(&users)
	if(result.Error != nil){
		e := HTTPError{Message:result.Error.Error(),Code:500}
		return nil, &e
	}
	return users,nil

}

func CreateUser(userDto dtos.CreateUserDto) (*HTTPError) {
	var user User
	result := db.DBConn.Where(&User{Username:userDto.Username}).First(&user)
	if(result.RowsAffected != 0){
		e := HTTPError{Message:"This username has already exist",Code:400}
		return &e
	}
	//Hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password),bcrypt.DefaultCost);
	if(err != nil){
		panic(err);
	}
	user = User{
		Username: userDto.Username,
		Password: string(hashedPassword),
	}
	result = db.DBConn.Create(&user)
	if(result.Error != nil){
		e := HTTPError{Message:result.Error.Error(),Code:500}
		return &e
	}
	return nil;
}