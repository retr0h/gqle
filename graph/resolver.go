package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"go-graphql-api/config"
	"gorm.io/gorm"
)

type Resolver struct {
	Database *gorm.DB
	Config   *config.Config
}
