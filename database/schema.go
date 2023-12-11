package database

import (
	"RBAC/config"
	"fmt"
)

func CreateTables() {
	dbObject, err := DB.DB()
	if err != nil {
		fmt.Println(err)
	}

	dbObject.Exec(fmt.Sprintf("USE %s", config.Config("DB_NAME")))

	// Create table for Products
	dbObject.Exec(`
	CREATE TABLE IF NOT EXISTS products (
		id int NOT NULL AUTO_INCREMENT,
		name varchar(255),
		description varchar(255),
		category varchar(255),
		amount int,
		PRIMARY KEY (id)
		);`,
	)

	// Users
	dbObject.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id int NOT NULL AUTO_INCREMENT,
		full_name varchar(255),
		email varchar(255),
		password varchar(255),
		PRIMARY KEY (id)
		);`,
	)

	// Groups
	dbObject.Exec(`
	CREATE TABLE IF NOT EXISTS groups (
		id int NOT NULL AUTO_INCREMENT,
		name varchar(255),
		description varchar(255),
		permissions varchar(255),
		user_id int,
		PRIMARY KEY (id)
		);`,
	)
}
