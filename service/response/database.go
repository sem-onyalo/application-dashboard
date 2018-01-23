package response

import "github.com/jinzhu/gorm"

// NewConnection represents the response to a new connection request
type NewConnection struct {
	Store *gorm.DB
}
