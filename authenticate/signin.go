package authenticate

import (
	"encoding/json"
	"log"
	"net/http"

	config "../config"
	encrypt "../encrypt"
	response "../model"
)

type DriverAccount struct {
	driverid string `json:"driverid"`
	email    string `json:email`
	password string `json:password`
}

type Driver struct {
	driverid string `json:"driverid"`
	email    string `json:email`
}

type DetailDriver struct {
	iddriver    string `json:"iddriver"`
	emaildriver string `json:emaildriver`
	// driverID string `json:"driverID"`
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var response response.Token
	var params = map[string]string{}

	err := r.ParseMultipartForm(4096)

	if err != nil {
		panic(err)
	}

	params["email"] = r.FormValue("email")
	params["password"] = r.FormValue("password")

	tokenStringsss, err := ValidateUser(params)
	idDriver, err := GetDataDriver(params)
	if tokenStringsss == "Incorrect email or password" {
		log.Print(err)
	} else {
		log.Print("token", tokenStringsss)
		response.Token = tokenStringsss
		response.DriverID = idDriver.iddriver
		response.DriverEmail = idDriver.emaildriver
		response.TokenType = "bearer"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ValidateUser(params map[string]string) (string, error) {
	var tokenString string
	var driver DriverAccount

	db := config.Connect()
	defer db.Close()

	email := params["email"]

	rows, err := db.Query("SELECT driverid, driver_email, driver_password FROM driver WHERE driver_email = $1", email)
	log.Print("ini rows", err)
	if err != nil {
		return "", err
	}

	for rows.Next() {
		if err := rows.Scan(&driver.driverid, &driver.email, &driver.password); err != nil {

			panic(err.Error())
		} else {
			log.Print("ini driver", driver)
			payload := driver

			driverID := payload.driverid
			email := payload.email
			password := payload.password
			pass := params["password"]

			verify := encrypt.ComparePasswords(password, []byte(pass))

			if verify == false {
				return "Incorrect email or password", nil
			} else {
				var params map[string]string
				params = map[string]string{}

				params["driverID"] = driverID
				params["email"] = email

				getToken, err := GenerateToken(params)

				if err != nil {
					return "", err
				}

				log.Print("ini nihh", driverID)

				tokenString := getToken

				return tokenString, nil
			}
		}
	}

	return tokenString, nil
}

func GetDataDriver(params map[string]string) (DetailDriver, error) {
	// var DataDriver string
	var DataDriver DetailDriver
	var driver Driver
	// var detail DetailDriver

	db := config.Connect()
	defer db.Close()

	email := params["email"]

	rows, err := db.Query("SELECT driverid, driver_email FROM driver WHERE driver_email = $1", email)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&driver.driverid, &driver.email); err != nil {
			panic(err.Error())
		} else {

			DataDriver.iddriver = driver.driverid
			DataDriver.emaildriver = driver.email

		}
	}

	// DataDriver
	log.Print("detail", DataDriver)
	return DataDriver, nil

}
