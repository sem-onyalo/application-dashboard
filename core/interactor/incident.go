package interactor

import (
	"errors"

	"github.com/sem-onyalo/application-dashboard/core/entity"
	"github.com/sem-onyalo/application-dashboard/service"
	"github.com/sem-onyalo/application-dashboard/service/request"
	"github.com/sem-onyalo/application-dashboard/service/response"
)

// TODO: move to config service
const (
	defaultIncidentResolutionOrder = "created_at desc"
	selectIncidentResolutions      = "incident.id, resolution.id, resolution.created_at, resolution.name, resolution.details"
)

// Incident represents an interactor to incidents
type Incident struct {
	Database   service.Database
	Resolution service.Resolution
}

// NewIncident creates a reference to an incident interactor
func NewIncident(database service.Database, resolution service.Resolution) *Incident {
	return &Incident{Database: database, Resolution: resolution}
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

// CreateResolution creates a new incident resolution
func (i Incident) CreateResolution(req request.CreateIncidentResolution) (response.CreateIncidentResolution, error) {
	var res response.CreateIncidentResolution

	conn, err := i.Database.NewConnection()
	if err != nil {
		return res, err
	}
	defer conn.Store.Close()

	var incident entity.Incident
	conn.Store.First(&incident, req.IncidentID)
	if incident.ID <= 0 {
		return res, errors.New("Invalid incident ID")
	}

	createResolution, err := i.Resolution.Create(request.CreateResolution{
		Name:    req.Name,
		Details: req.Details,
	})
	if err != nil {
		return res, err
	}

	incidentResolution := entity.IncidentResolution{
		IncidentID:   incident.ID,
		ResolutionID: createResolution.ID,
	}
	conn.Store.Create(&incidentResolution)

	res.IncidentID = req.IncidentID
	res.ResolutionName = req.Name
	res.ResolutionDetails = req.Details

	return res, nil
}

// GetAllResolutions gets all incident resolutions
func (i Incident) GetAllResolutions() (response.GetAllIncidentResolutions, error) {
	var res response.GetAllIncidentResolutions

	conn, err := i.Database.NewConnection()
	if err != nil {
		return res, err
	}
	defer conn.Store.Close()

	var resolutions []entity.IncidentResolution
	rows, err := conn.Store.
		Table("resolution").
		Order(defaultIncidentResolutionOrder).
		Limit(defaultSelectLimit).
		Select(selectIncidentResolutions).
		Joins("inner join incident_resolution on incident_resolution.resolution_id = resolution.id").
		Joins("inner join incident on incident.id = incident_resolution.incident_id").
		Rows()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var resolution entity.IncidentResolution
		err = rows.Scan(
			&resolution.IncidentID,
			&resolution.ResolutionID,
			&resolution.ResolutionCreatedAt,
			&resolution.ResolutionName,
			&resolution.ResolutionDetails,
		)
		if err != nil {
			return res, err
		}

		resolutions = append(resolutions, resolution)
	}

	res.Resolutions = resolutions
	return res, nil
}
