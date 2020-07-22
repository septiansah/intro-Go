package main

import (
	"log"
	"net/http"

	authenticate "./authenticate"
	endpoint "./endpoint"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/signin", authenticate.Authenticate).Methods("POST")
	router.HandleFunc("/tokenInfo/{token}", authenticate.ValidateToken).Methods("GET")

	router.HandleFunc("/getDrivers", endpoint.ReturnAllDriver).Methods("GET")
	router.HandleFunc("/getDriver/{driverid}", endpoint.GetDriver).Methods("GET")
	router.HandleFunc("/updateDriver/{driverid}", endpoint.UpdateDriver).Methods("PUT")
	router.HandleFunc("/getVehicles", endpoint.GetVehicles).Methods("GET")
	// router.HandleFunc("/createDriver", createDriver).Methods("POST")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":5050", router))

}
