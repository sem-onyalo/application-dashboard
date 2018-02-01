package entity

import "time"

// IncidentResolution represents an incident IncidentResolution
type IncidentResolution struct {
	IncidentID          uint
	ResolutionID        uint
	ResolutionCreatedAt time.Time `gorm:"-"`
	ResolutionName      string    `gorm:"-"`
	ResolutionDetails   string    `gorm:"-"`
}
