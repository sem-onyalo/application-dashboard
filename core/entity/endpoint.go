package entity

import (
	"github.com/jinzhu/gorm"
)

// Endpoint is a representation of an API web endpoint
type Endpoint struct {
	gorm.Model
	Name string
	URL  string
}
