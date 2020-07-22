package endpoint

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	model "../model"

	config "../config"

	"github.com/gorilla/mux"
)

func ReturnAllDriver(w http.ResponseWriter, r *http.Request) {
	var drivers model.Drivers
	var arr_drivers []model.Drivers
	var responseDrivers model.ResponseDrivers

	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("Select driverid, driver_fullname, driver_email from driver")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&drivers.Driverid, &drivers.Driver_fullname, &drivers.Driver_email); err != nil {
			log.Fatal(err.Error())

		} else {
			arr_drivers = append(arr_drivers, drivers)
		}
	}

	responseDrivers.Status = 1
	responseDrivers.Messages = "Success"
	responseDrivers.Data = arr_drivers

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseDrivers)

}

func GetDriver(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Method", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	driverID := mux.Vars(r)["driverid"]

	var driver model.Driver
	var arrDriver []model.Driver
	// var responseDriver model.CombineRes
	log.Print("bro", &arrDriver)
	var vehicle model.Vehicle
	// var arrVehicle []model.Vehicle
	var response model.ResponseDriver

	db := config.Connect()
	defer db.Close()

	query := fmt.Sprintf("SELECT driverid, driver_citizen_numberid, driver_company_transporterid, driver_created_on, driver_email, COALESCE(driver_fullname, ''), driver_internalid, driver_last_update_on, driver_license_numberid, driver_phone_number, driver_status_active, driver_status, COALESCE(driver_photo, '')  FROM driver WHERE driverid = '%s'", driverID)
	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {

		if err := rows.Scan(
			&driver.DriverID,
			&driver.DriverCitizenID,
			&driver.DriverCompanyTransporterID,
			&driver.DriverCreatedOn,
			&driver.DriverEmail,
			&driver.DriverFullname,
			&driver.DriverInternalID,
			&driver.DriverLastUpdate,
			&driver.DriverLicenseNumberID,
			&driver.DriverPhoneNumber,
			&driver.DriverStatusActive,
			&driver.DriverStatus,
			&driver.DriverPhoto); err != nil {
			log.Fatal(err.Error())
		} else {
			db1 := config.ConnectVehicle()
			defer db1.Close()

			query1 := fmt.Sprintf("SELECT vehicle_transporterid, vehicle_company_transporterid, vehicle_dedicated_for_company_shipperid, vehicle_transporterkirnumber, vehicle_transporter_maximumcbmdimension, vehicle_transporter_maximum_weight, vehicle_transporter_plate_number FROM vehicle_transporter WHERE vehicle_transporterid = 'TK-VHC-201912131918140000001'")
			rows1, err1 := db1.Query(query1)

			if err1 != nil {
				log.Print(err1)
			}

			for rows1.Next() {
				if err1 := rows1.Scan(&vehicle.VehicleID,
					&vehicle.VehicleCompanyTransporterID,
					&vehicle.VehicleDedicatedForCompanyShipperid,
					&vehicle.VehicleTransporterkirnumber,
					&vehicle.VehicleTransporterMaximumcbmdimension,
					&vehicle.VehicleTransporterMaximumWeight,
					&vehicle.VehicleTransportePlateNumber); err1 != nil {
					log.Fatal(err1.Error())
				} else {

					driver.Tes = append(driver.Tes, vehicle)
				}
			}

			arrDriver = append(arrDriver, driver)

		}
	}

	response.Status = 1
	response.Messages = "Success"
	response.Data = arrDriver

	log.Print("[GET] DATA USER")

	w.Header().Set("Conten-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func createDriver(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Method", "POST, GET, OPTION, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var response model.ResponseCreate

	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	driverid := r.FormValue("driverId")
	driverFullname := r.FormValue("driverFullname")
	driverEmail := r.FormValue("driverEmail")
	date := time.Now().Format(time.RFC3339)

	query := fmt.Sprintf("INSERT INTO driver (driverid, driver_fullname, driver_email, driver_created_on) VALUES('%s', '%s', '%s', '%s')", driverid, driverFullname, driverEmail, date)

	_, err = db.Exec(query)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Messages = "Success"

	log.Print("[POST] INSERT DRIVER")

	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateDriver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	driverID := mux.Vars(r)["driverid"]

	var response model.ResponseCreate

	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	driverCitizenID := r.FormValue("driver_citizenID")
	driverEmail := r.FormValue("driver_email")
	driverFullname := r.FormValue("driver_fullname")
	driverPhoneNumber := r.FormValue("driver_phonenumber")
	log.Print(driverID)
	query := fmt.Sprintf("UPDATE driver SET driver_fullname = '%s', driver_email = '%s', driver_citizen_numberid = '%s', driver_phone_number = '%s' WHERE driverid = '%s' ",
		driverFullname,
		driverEmail,
		driverCitizenID,
		driverPhoneNumber,
		driverID,
	)

	_, err = db.Exec(query)

	if err != nil {
		response.Status = 0
		response.Messages = "Failed"

		log.Print(err)
	} else {
		response.Status = 1
		response.Messages = "Success"
		log.Print("[UPDATE] UPDATE DRIVER DATA TO DATABASE")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func GetVehicles(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Method", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// driverID := mux.Vars(r)["driverid"]

	var vehicle model.Vehicle
	var arrVehicle []model.Vehicle
	var responseVehicle model.ResponseVehicles

	db := config.ConnectVehicle()
	defer db.Close()

	query := fmt.Sprintf("SELECT vehicle_transporterid, vehicle_company_transporterid, vehicle_dedicated_for_company_shipperid, vehicle_transporterkirnumber, vehicle_transporter_maximumcbmdimension, vehicle_transporter_maximum_weight, vehicle_transporter_plate_number FROM vehicle_transporter")
	rows, err := db.Query(query)

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&vehicle.VehicleID, &vehicle.VehicleCompanyTransporterID, &vehicle.VehicleDedicatedForCompanyShipperid, &vehicle.VehicleTransporterkirnumber, &vehicle.VehicleTransporterMaximumcbmdimension, &vehicle.VehicleTransporterMaximumWeight, &vehicle.VehicleTransportePlateNumber); err != nil {
			log.Fatal(err.Error())
		} else {
			arrVehicle = append(arrVehicle, vehicle)
		}
	}
	log.Print(arrVehicle)
	responseVehicle.Status = 1
	responseVehicle.Messages = "Success"
	responseVehicle.Data = arrVehicle
	log.Print("[GET] DATA USER")

	w.Header().Set("Conten-Type", "application/json")
	json.NewEncoder(w).Encode(responseVehicle)

}
