package database

import (
	"fmt"

	"go-graphql-api/dbmodel"
	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DBInstance *gorm.DB
	err        error
)

// connecting to the db
func ConnectDB() {
	DBInstance, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
		panic("Database connection attempt was unsuccessful.....")
	} else {
		fmt.Println("Database Connected successfully.....")
	}
}

func CreateDB() {
	// Create a database
	DBInstance.Exec("CREATE DATABASE IF NOT EXISTS Blog_Posts")
	// make the database available to this connection
	DBInstance.Exec("USE Blog_Posts")
}

func MigrateDB() {
	// migrate and sync the model to create a db table
	DBInstance.AutoMigrate(&dbmodel.Post{})
	fmt.Println("Database migration completed....")
}
