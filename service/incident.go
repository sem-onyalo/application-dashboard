package service

import (
	"github.com/sem-onyalo/application-dashboard/service/request"
	"github.com/sem-onyalo/application-dashboard/service/response"
)

// Incident is a boundary to interacting with incidents
type Incident interface {
	Create(request.CreateIncident) (response.CreateIncident, error)
}
