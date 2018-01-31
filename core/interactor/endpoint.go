package interactor

import (
	"errors"
	"net/http"
	"time"

	"github.com/sem-onyalo/application-dashboard/core/entity"
	"github.com/sem-onyalo/application-dashboard/service"
	"github.com/sem-onyalo/application-dashboard/service/request"
	"github.com/sem-onyalo/application-dashboard/service/response"
)

// TODO: move to config service
const (
	defaultEndpointTestLimit  = 1000
	defaultEndpointTestOrder  = "created_at desc"
	defaultEndpointTestResult = "NULL"
	selectEndpointTests       = "endpoint.id, endpoint.name, endpoint.url, endpoint_test.id, endpoint_test.created_at, endpoint_test.updated_at, endpoint_test.response_status, endpoint_test.time_elapsed"
)

// Endpoint is an interactor for interacting with the endpoint entity
type Endpoint struct {
	Database service.Database
	Incident service.Incident
}

// NewEndpoint creates a pointer to an endpoint interactor
func NewEndpoint(database service.Database, incident service.Incident) *Endpoint {
	return &Endpoint{Database: database, Incident: incident}
}

// GetAll function gets all endpoints from the datastore
func (e Endpoint) GetAll() (response.GetAllEndpoints, error) {
	var r response.GetAllEndpoints

	conn, err := e.Database.NewConnection()
	if err != nil {
		return r, err
	}
	defer conn.Store.Close()

	var endpoints []entity.Endpoint
	conn.Store.Find(&endpoints)
	r = response.GetAllEndpoints{Endpoints: endpoints}
	return r, nil
}

// TestAll function tests all endpoints in the datastore
func (e Endpoint) TestAll() (response.TestAllEndpoints, error) {
	var r response.TestAllEndpoints
	getAllEndpoints, err := e.GetAll()
	if err != nil {
		return r, err
	}

	var endpointTests []entity.EndpointTest
	for _, ep := range getAllEndpoints.Endpoints {
		timerStart := time.Now()
		// TODO: move this to an http service
		rsp, err := http.Get(ep.URL)
		timerEnd := time.Now()
		timerElapsed := timerEnd.Sub(timerStart)

		var result string
		if err != nil {
			result = defaultEndpointTestResult
		} else {
			result = rsp.Status
		}

		endpointTest := entity.EndpointTest{
			EndpointID:     ep.ID,
			Name:           ep.Name,
			URL:            ep.URL,
			ResponseStatus: result,
			TimeElapsed:    timerElapsed.Seconds(),
		}
		err = e.saveTest(&endpointTest)
		if err != nil {
			return r, err
		}
		endpointTests = append(endpointTests, endpointTest)
	}

	r = response.TestAllEndpoints{EndpointTests: endpointTests}
	return r, nil
}

// GetTests retrieves a collection of endpoint test records
func (e Endpoint) GetTests() (response.GetEndpointTests, error) {
	var r response.GetEndpointTests
	conn, err := e.Database.NewConnection()
	if err != nil {
		return r, err
	}
	defer conn.Store.Close()

	var endpointTests []entity.EndpointTest
	rows, err := conn.Store.
		Table("endpoint_test").
		Order(defaultEndpointTestOrder).
		Limit(defaultEndpointTestLimit).
		Select(selectEndpointTests).
		Joins("inner join endpoint on endpoint.id = endpoint_test.endpoint_id").
		Rows()
	if err != nil {
		return r, err
	}

	for rows.Next() {
		var endpointTest entity.EndpointTest
		err = rows.Scan(&endpointTest.EndpointID,
			&endpointTest.Name,
			&endpointTest.URL,
			&endpointTest.ID,
			&endpointTest.CreatedAt,
			&endpointTest.UpdatedAt,
			&endpointTest.ResponseStatus,
			&endpointTest.TimeElapsed,
		)
		if err != nil {
			return r, err
		}

		endpointTests = append(endpointTests, endpointTest)
	}

	r = response.GetEndpointTests{EndpointTests: endpointTests}
	return r, nil
}

// CreateIncident creates a new incident for an endpoint
func (e Endpoint) CreateIncident(req request.CreateEndpointIncident) (response.CreateEndpointIncident, error) {
	var res response.CreateEndpointIncident

	conn, err := e.Database.NewConnection()
	if err != nil {
		return res, err
	}
	defer conn.Store.Close()

	var endpoint entity.Endpoint
	conn.Store.First(&endpoint, req.EndpointID)
	if endpoint.ID <= 0 {
		return res, errors.New("Invalid endpoint ID")
	}

	createIncident, err := e.Incident.Create(request.CreateIncident{
		Urgency: req.Urgency,
		Impact:  req.Impact,
		Details: req.Details,
	})
	if err != nil {
		return res, err
	}

	var endpointIncident = entity.EndpointIncident{
		EndpointID: req.EndpointID,
		IncidentID: createIncident.ID,
	}
	conn.Store.Create(&endpointIncident)

	res.EndpointID = req.EndpointID
	res.EndpointName = endpoint.Name
	res.IncidentUrgency = req.Urgency
	res.IncidentImpact = req.Impact
	res.IncidentDetails = req.Details
	return res, nil
}

// saveTest saves an endpoint test record
func (e Endpoint) saveTest(endpointTest *entity.EndpointTest) error {
	conn, err := e.Database.NewConnection()
	if err != nil {
		return err
	}
	defer conn.Store.Close()

	conn.Store.Create(endpointTest)
	return nil
}
