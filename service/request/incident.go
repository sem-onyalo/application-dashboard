package request

// CreateIncident represents a request to create an incident
type CreateIncident struct {
	Urgency int
	Impact  int
	Details string
}
