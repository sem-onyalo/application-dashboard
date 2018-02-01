package response

import "github.com/sem-onyalo/application-dashboard/core/entity"

// CreateIncident represents a response to a create incident request
type CreateIncident struct {
	ID uint
}

// CreateIncidentResolution represents a response to a create incident resolution request
type CreateIncidentResolution struct {
	IncidentID        uint
	ResolutionName    string
	ResolutionDetails string
}

// GetAllIncidentResolutions represents a response to a get all incident resolution requests
type GetAllIncidentResolutions struct {
	Resolutions []entity.IncidentResolution
}
