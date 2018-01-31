package entity

import "time"

// EndpointIncident represents an endpoint incident
type EndpointIncident struct {
	EndpointID        uint
	IncidentID        uint
	IncidentCreatedAt time.Time `gorm:"-"`
	EndpointName      string    `gorm:"-"`
	IncidentUrgency   int       `gorm:"-"`
	IncidentImpact    int       `gorm:"-"`
	IncidentDetails   string    `gorm:"-"`
}
