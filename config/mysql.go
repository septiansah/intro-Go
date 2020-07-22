package config

import (
	"database/sql"
	"fmt"
	"log"
)

func Connect() *sql.DB {
	const (
		hostname     = "localhost"
		host_port    = 5432
		username     = "postgres"
		password     = "admin123"
		databasename = "belajarGo"
	)

	pg_con_string := fmt.Sprintf("port=%d host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host_port, hostname, username, password, databasename)

	db, err := sql.Open("postgres", pg_con_string)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
func ConnectVehicle() *sql.DB {
	const (
		hostname     = "localhost"
		host_port    = 5432
		username     = "postgres"
		password     = "admin123"
		databasename = "vehicle"
	)

	pg_con_string := fmt.Sprintf("port=%d host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host_port, hostname, username, password, databasename)

	db, err := sql.Open("postgres", pg_con_string)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
