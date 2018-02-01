package request

// CreateIncident represents a request to create an incident
type CreateIncident struct {
	Urgency int
	Impact  int
	Details string
}

// CreateIncidentResolution represents a request to create an incident resolution
type CreateIncidentResolution struct {
	IncidentID uint
	Name       string
	Details    string
}
