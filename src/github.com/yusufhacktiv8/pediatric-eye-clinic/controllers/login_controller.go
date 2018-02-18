package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/yusufhacktiv8/pediatric-eye-clinic/models"
)

type LoginController struct {
	DB *sql.DB
}

func (a *LoginController) Authenticate(w http.ResponseWriter, r *http.Request) {

	jwtSecretKey := "jshbdgh54gs9jdbx543GnhY67"

	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	errFindUser := user.FindOne(a.DB)
	if errFindUser != nil {
		result := map[string]interface{}{"status": "LOGIN_ERROR"}
		respondWithJSON(w, http.StatusCreated, result)
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email":    user.Email,
			"password": user.Password,
		})
		tokenString, error := token.SignedString([]byte(jwtSecretKey))
		if error != nil {
			fmt.Println(error)
		}
		result := map[string]interface{}{"token": tokenString}
		respondWithJSON(w, http.StatusCreated, result)
	}
}
