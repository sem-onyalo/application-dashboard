package entity

import "github.com/jinzhu/gorm"

// EndpointTest is a representation of an API web endpoint test
type EndpointTest struct {
	gorm.Model
	Name           string `gorm:"-"`
	URL            string `gorm:"-"`
	EndpointID     uint
	ResponseStatus string
	TimeElapsed    float64
}
