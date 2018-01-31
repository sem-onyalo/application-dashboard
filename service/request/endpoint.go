package request

// CreateEndpointIncident represents a request to create an endpoint incident
type CreateEndpointIncident struct {
	EndpointID uint
	Urgency    int
	Impact     int
	Details    string
}
