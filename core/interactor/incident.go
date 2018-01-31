package interactor

import (
	"github.com/sem-onyalo/application-dashboard/core/entity"
	"github.com/sem-onyalo/application-dashboard/service"
	"github.com/sem-onyalo/application-dashboard/service/request"
	"github.com/sem-onyalo/application-dashboard/service/response"
)

// Incident represents an interactor to incidents
type Incident struct {
	Database service.Database
}

// NewIncident creates a reference to an incident interactor
func NewIncident(database service.Database) *Incident {
	return &Incident{Database: database}
}

// Create creates a new incident
func (i Incident) Create(request request.CreateIncident) (response.CreateIncident, error) {
	var r response.CreateIncident

	conn, err := i.Database.NewConnection()
	if err != nil {
		return r, err
	}
	defer conn.Store.Close()

	var incident = entity.Incident{
		Urgency: request.Urgency,
		Impact:  request.Impact,
		Details: request.Details,
	}

	conn.Store.Create(&incident)
	r.ID = incident.ID
	return r, nil
}
