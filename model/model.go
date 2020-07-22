package model

import (
	_ "github.com/lib/pq"
)

type Drivers struct {
	Driverid        string `form:"driverid" json:"driverid"`
	Driver_fullname string `form:"driver_fullname" json:"driverFullname"`
	Driver_email    string `form:"driver_email" json:"driver_email"`
}

type ResponseDrivers struct {
	Status   int       `json:"status"`
	Messages string    `json:"messages"`
	Data     []Drivers `json:"data"`
}

type ResponseCreate struct {
	Status   int    `json:"status"`
	Messages string `json:"messages"`
}

type Driver struct {
	DriverID                   string    `form:"driverid" json:"driverid"`
	DriverCitizenID            string    `form:"driver_citizen_numberid" json:"driver_citizen_numberid"`
	DriverCompanyTransporterID string    `form:"driver_company_transporterid" json:"driver_company_transporterid"`
	DriverCreatedOn            string    `form:"driver_created_on" json:"driver_created_on"`
	DriverEmail                string    `form:"driver_email" json:"driver_email"`
	DriverFullname             string    `form:"driver_fullname" json:"driver_fullname"`
	DriverInternalID           string    `form:"driver_internaid" json:"driver_internaid"`
	DriverLastUpdate           string    `form:"driver_last_update" json:"driver_last_update"`
	DriverLicenseNumberID      string    `form:"driver_license_numberid" json:"driver_license_numberid"`
	DriverPhoneNumber          string    `form:"driver_phone_number" json:"driver_phone_number"`
	DriverStatusActive         string    `form:"driver_status_active" json:"driver_status_active"`
	DriverStatus               bool      `form:"driver_status" json:"driver_status"`
	DriverPhoto                string    `form:"driver_photo" json:"driver_photo"`
	Tes                        []Vehicle `json:"VehicleDetail"`
}

type CombineRes struct {
	DriverRes  []Driver  `json:"data"`
	VehicleRes []Vehicle `json:"data1"`
}

type ResponseDriver struct {
	Status   int      `json:"status"`
	Messages string   `json:"messages"`
	Data     []Driver `json:"data"`
	// Tes      string   `json:"data"`
}

type Vehicle struct {
	VehicleID                             string `form:"vehicleId", json:"vehicleId"`
	VehicleCompanyTransporterID           string `form:"vehicle_transporterid", json:"vehicle_transporterid"`
	VehicleDedicatedForCompanyShipperid   string `form:"vehicle_dedicated_for_company_shipperid", json:"vehicle_dedicated_for_company_shipperid"`
	VehicleTransporterkirnumber           string `form:"vehicle_transporterkirnumber", json:"vehicle_transporterkirnumber"`
	VehicleTransporterMaximumcbmdimension string `form:"vehicle_transporter_maximumcbmdimension", json:"vehicle_transporter_maximumcbmdimension"`
	VehicleTransporterMaximumWeight       string `form:"vehicle_transporter_maximum_weight", json:"vehicle_transporter_maximum_weight"`
	VehicleTransportePlateNumber          string `form:"vehicle_transporter_plate_number1212", json:"vehicle_transporter_plate_number1212"`
}

type ResponseVehicles struct {
	Status   int       `json:"status"`
	Messages string    `json:"messages"`
	Data     []Vehicle `json:"data"`
}

type Token struct {
	Token       string `json:"token"`
	TokenType   string `json:"token_type"`
	DriverID    string `json:"driverID"`
	DriverEmail string `json:"DriverEmail"`
}

type TokenInfo struct {
	DriverID    string `json:"driverID"`
	DriverEmail string `json:"driverEmail"`
	Hit         string `json:"hit"`
}
