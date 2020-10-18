package auth

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

//New ...
func New(w http.ResponseWriter, r *http.Request) {

	type Username struct {
		Name string `json:"key"`
	}

	type Token struct {
		Name string `json:"accessToken"`
	}

	var username Username

	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token := generateToken(username.Name)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var data Token

	data.Name = token

	res, err := json.Marshal(&data)

	print(res, token)

	w.Write(res)

}

func generateToken(username string) string {

	type customClaims struct {
		Username string
		jwt.StandardClaims
	}

	claims := customClaims{
		Username:       username,
		StandardClaims: jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte("secure"))

	if err != nil {
		print(err)
	}

	return signedToken
}
