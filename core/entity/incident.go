package entity

import "github.com/jinzhu/gorm"

// Incident represents an incident
type Incident struct {
	gorm.Model
	Urgency int
	Impact  int
	Details string
}
