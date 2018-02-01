package entity

import (
	"github.com/jinzhu/gorm"
)

// Resolution represents a resolution
type Resolution struct {
	gorm.Model
	Name    string
	Details string
}
