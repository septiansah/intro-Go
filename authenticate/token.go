package authenticate

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	// jwt "../github.com/dgrijalva/jwt-go"
	response "../model"
	"github.com/gorilla/mux"
)

type Execption struct {
	Messages string `form:"messages"`
}

var secretKey = []byte("truckking")

func GenerateToken(params map[string]string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["driverID"] = params["driverID"]
	claims["email"] = params["email"]
	claims["hit"] = time.Now().Unix()

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err

	}

	return tokenString, nil
}

func ValidateToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	var response response.TokenInfo

	token := mux.Vars(r)["token"]

	info, err := GetTokenInfo(token)

	if err != nil {
		json.NewEncoder(w).Encode(Execption{Messages: "Invalid authorization token"})
	} else {
		response.DriverID = info.DriverID
		response.DriverEmail = info.DriverEmail
		response.Hit = info.Hit

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func GetTokenInfo(tokenString string) (response.TokenInfo, error) {
	var TokenInfo response.TokenInfo

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err == nil {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			userID := claims["driverID"]
			email := claims["email"]
			Hit := claims["Hit"].(float64)

			t := time.Unix(int64(Hit), 0)

			TokenInfo.DriverID = userID.(string)
			TokenInfo.DriverEmail = email.(string)
			TokenInfo.Hit = t.Format(time.RFC3339)

			return TokenInfo, nil
		} else {
			return TokenInfo, nil
		}
	} else {
		return TokenInfo, err
	}
}

func RandomString() (string, error) {
	b := make([]byte, 3)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
