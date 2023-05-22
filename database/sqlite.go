package database

import (
	"github.com/retr0h/gqle/config"
	"github.com/retr0h/gqle/dbmodel"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	sqllogger "gorm.io/gorm/logger"
)

// Database handles database access
type Database struct {
	config   *config.Config
	logger   *logrus.Logger
	Database *gorm.DB
}

// Connect connect to the database and populate the structs Database field
func (db *Database) Connect() {
	gormConfig := &gorm.Config{}
	if db.config.Debug {
		gormConfig.Logger = sqllogger.Default.LogMode(sqllogger.Info)
	}

	instance, err := gorm.Open(sqlite.Open(db.config.DSN), gormConfig)
	if err != nil {
		db.logger.Fatalf("Database connection failed: %s", err)
	} else {
		db.Database = instance
		db.logger.WithFields(logrus.Fields{
			"DBName": db.config.DBName,
			"DSN":    db.config.DSN,
		}).Info("Database connected successfully")
	}
}

// Migrate migrate the database
func (db *Database) Migrate() {
	db.Database.AutoMigrate(&dbmodel.Post{})
	db.Database.AutoMigrate(&dbmodel.VPC{})
	db.logger.WithFields(logrus.Fields{
		"DBName": db.config.DBName,
		"DSN":    db.config.DSN,
	}).Info("Database migration completed")
}

// New create a new database instance
func New(
	config *config.Config,
	logger *logrus.Logger,
) *Database {
	return &Database{
		config: config,
		logger: logger,
	}
}
