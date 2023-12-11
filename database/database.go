package database

import (
	"RBAC/config"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database instance
var DB *gorm.DB

// Connect function
func Connect() error {
	var err error

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=%s&parseTime=%s&loc=%s",
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		"tcp",
		config.Config("DB_HOST"),
		config.Config("DB_NAME"),
		"utf8",
		"True",
		"Local",
	)

	log.Print(">> dsn: ", dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return err
	}

	dbObject, err := DB.DB()
	if err != nil {
		return err
	}

	if err = dbObject.Ping(); err != nil {
		return errors.New("ping database error")
	}

	// From schema.go, create tables
	CreateTables()

	return nil
}
