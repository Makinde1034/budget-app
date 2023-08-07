package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/Makinde1034/budget-app/model"
	"golang.org/x/crypto/bcrypt"

	"github.com/Makinde1034/budget-app/helpers"
	"log"

	uuid "github.com/satori/go.uuid"
)

func hashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14) 
      
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

func verifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword),[]byte(providedPassword))
	check := true
	msg := ""

	if err != nil {
		check = false
		msg = "Incorrect password"
	}

	return check, msg
}

func RegisterUser(w http.ResponseWriter, r *http.Request){
	createUserRequest := model.User{}

	json.NewDecoder(r.Body).Decode(&createUserRequest)

	userExists,_ := model.FetchUserByEmail(createUserRequest.Email)

	if len(userExists) > 0 {
		err := model.ErrMsg{Msg: "User with email already exists"}
		
		w.WriteHeader(http.StatusConflict)   
		json.NewEncoder(w).Encode(err)
		return
	}

	hashedPassword := hashPassword(createUserRequest.Password)       
	emailToken := uuid.NewV4()

	user := model.User{
		Email: createUserRequest.Email,
		Password:  hashedPassword,
		EmailVerificationToken:emailToken,
	}

	err,_ := model.CreateUser(&user)

	if err != nil{
		errMsg := model.ErrMsg{Msg:  "Request failed. Please try again later"}

		json.NewEncoder(w).Encode(errMsg)
		return
	}

	fetchUser,_ := model.FetchUserByEmail(createUserRequest.Email)

	token, _ := helpers.GenerateToken(createUserRequest.Email,fetchUser[0].ID)

	response := struct{
		User model.User `json:"user"`
		Token string `json:"token"`
	}{
		fetchUser[0],
		token,
	}

	json.NewEncoder(w).Encode(response) 

}

func Login(w http.ResponseWriter, r *http.Request){
	loginRequest := model.User{}

	json.NewDecoder(r.Body).Decode(&loginRequest)

	userExists,_ := model.FetchUserByEmail(loginRequest.Email)      

	if len(userExists) == 0 {
		errMsg := model.ErrMsg{Msg:  "Check your email and try again"}
		w.WriteHeader(http.StatusConflict)   
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	check, msg := verifyPassword(userExists[0].Password,loginRequest.Password)


	if !check {
		errMsg := struct{                                                  
			Msg string `json:"msg"`
		}{msg}

		w.WriteHeader(http.StatusConflict)                               
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	token, _ := helpers.GenerateToken(loginRequest.Email,userExists[0].ID)

	response := struct{
		User model.User `json:"user"`
		Token string `json:"token"`
	}{
		userExists[0],
		token,
	}
	json.NewEncoder(w).Encode(response) 

	
}

func ValidateToken(w http.ResponseWriter, r *http.Request){             
	isTokenValid := true
	validateTokenRequest := struct{
		Token string `json:"token"`                                  
	}{
                                                                                
	}

	json.NewDecoder(r.Body).Decode(&validateTokenRequest)

	_, msg := helpers.VerifyToken(validateTokenRequest.Token)

	if msg != ""{
		isTokenValid = false
	}

	response := struct{
		IsTokenValid bool `json:"isTokenValid"`
	}{
		isTokenValid,
	}

	json.NewEncoder(w).Encode(response)


}