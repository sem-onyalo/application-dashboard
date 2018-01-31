package service

import (
	"github.com/sem-onyalo/application-dashboard/service/request"
	"github.com/sem-onyalo/application-dashboard/service/response"
)

// Endpoint is a boundary that allows you to interact with the endpoint entity
type Endpoint interface {
	GetAll() (response.GetAllEndpoints, error)
	TestAll() (response.TestAllEndpoints, error)
	GetTests() (response.GetEndpointTests, error)
	CreateIncident(request.CreateEndpointIncident) (response.CreateEndpointIncident, error)
}
