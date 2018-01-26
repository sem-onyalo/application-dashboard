package entity

// EndpointTest is a representation of an API web endpoint test
type EndpointTest struct {
	Name           string
	URL            string
	ResponseStatus string
	TimeElapsed    float64
}
