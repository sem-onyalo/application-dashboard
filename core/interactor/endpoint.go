package interactor

import (
	"net/http"
	"time"

	"github.com/sem-onyalo/application-dashboard/core/entity"
	"github.com/sem-onyalo/application-dashboard/service"
	"github.com/sem-onyalo/application-dashboard/service/response"
)

// Endpoint is an interactor for interacting with the endpoint entity
type Endpoint struct {
	Database service.Database
}

// NewEndpoint creates a pointer to an endpoint interactor
func NewEndpoint(database service.Database) *Endpoint {
	return &Endpoint{database}
}

// GetAll function gets all endpoints from the datastore
func (e Endpoint) GetAll() (response.GetAllEndpoints, error) {
	conn, err := e.Database.NewConnection()

	var r response.GetAllEndpoints
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
			result = "NULL"
		} else {
			result = rsp.Status
		}

		endpointTests = append(endpointTests, entity.EndpointTest{Name: ep.Name, URL: ep.URL, ResponseStatus: result, TimeElapsed: timerElapsed.Seconds()})
	}

	r = response.TestAllEndpoints{EndpointTests: endpointTests}
	return r, nil
}
