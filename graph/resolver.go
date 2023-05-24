package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"github.com/retr0h/gqle/pkg/config"
	"gorm.io/gorm"
)

type Resolver struct {
	Database *gorm.DB
	Config   *config.Config
}
