package response

import (
	"github.com/sem-onyalo/application-dashboard/core/entity"
)

// CreateEndpointIncident represents the response to a create endpoint incident request
type CreateEndpointIncident struct {
	EndpointID      uint
	EndpointName    string
	IncidentUrgency int
	IncidentImpact  int
	IncidentDetails string
}

// GetAllEndpoints represents the response to a GetAllEndpoints request
type GetAllEndpoints struct {
	Endpoints []entity.Endpoint
}

// GetAllEndpointIncidents represents the response to a GetAllEndpointIncidents request
type GetAllEndpointIncidents struct {
	Incidents []entity.EndpointIncident
}

// GetEndpointTests represents the response to get a collection of endpoint tests request
type GetEndpointTests struct {
	EndpointTests []entity.EndpointTest
}

// TestAllEndpoints represents the response to a TestAllEndpoints request
type TestAllEndpoints struct {
	EndpointTests []entity.EndpointTest
}
