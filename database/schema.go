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

	dbObject.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", config.Config("DB_NAME")))
	dbObject.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", config.Config("DB_NAME")))
	dbObject.Exec(fmt.Sprintf("USE %s", config.Config("DB_NAME")))

	// Create table for Products
	_, err = dbObject.Exec(`
	CREATE TABLE IF NOT EXISTS products (
		id int NOT NULL AUTO_INCREMENT,
		name varchar(255),
		description varchar(255),
		category varchar(255),
		amount int,
		PRIMARY KEY (id)
		);`)
	if err != nil {
		fmt.Println(err)
	}

	// Users
	_, err = dbObject.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id int NOT NULL AUTO_INCREMENT,
		full_name varchar(255),
		email varchar(255),
		password varchar(255),
		PRIMARY KEY (id)
		);`)
	if err != nil {
		fmt.Println(err)
	}

	// Groups
	_, err = dbObject.Exec(`
	CREATE TABLE IF NOT EXISTS user_groups (
		id int NOT NULL AUTO_INCREMENT,
		name varchar(255),
		description varchar(255),
		permissions varchar(255),
		user_id int,
		PRIMARY KEY (id)
		);`)
	if err != nil {
		fmt.Println(err)
	}
}

func InsertData() {
	dbObject, err := DB.DB()
	if err != nil {
		fmt.Println(err)
	}

	// Insert data into products table
	dbObject.Exec(`
        INSERT INTO products (name, description, category, amount) VALUES
        ("Apple", "iPhone 12", "Mobile", 100),
        ("Samsung", "Galaxy S21", "Mobile", 100),
        ("Samsung", "Galaxy S20", "Mobile", 100),
        ("Samsung", "Galaxy S10", "Mobile", 100),
        ("Samsung", "Galaxy S9", "Mobile", 100)
    `)

	// Insert data into users table
	dbObject.Exec(`
        INSERT INTO users (full_name, email, password) VALUES
        ("test", "test@test.com", "test"),
        ("admin", "admin@test.com", "admin")
    `)

	// Insert data into groups table
	dbObject.Exec(`
        INSERT INTO user_groups (name, description, permissions, user_id) VALUES
        ("admin", "Admin group", "read,write,delete", 1),
        ("user", "User group", "read", 2)
    `)
}
