package service

import (
	"github.com/sem-onyalo/application-dashboard/service/request"
	"github.com/sem-onyalo/application-dashboard/service/response"
)

// Endpoint is a boundary that allows you to interact with the endpoint entity
type Endpoint interface {
	CreateIncident(request.CreateEndpointIncident) (response.CreateEndpointIncident, error)
	GetAll() (response.GetAllEndpoints, error)
	GetAllIncidents() (response.GetAllEndpointIncidents, error)
	GetTests() (response.GetEndpointTests, error)
	TestAll() (response.TestAllEndpoints, error)
}
