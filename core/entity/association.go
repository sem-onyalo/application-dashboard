package entity

import (
	"github.com/jinzhu/gorm"
)

// Association is a representation of an association
type Association struct {
	gorm.Model
	Name string
}
