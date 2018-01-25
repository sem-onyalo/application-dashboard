package interactor

import (
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

	endpoints := make([]entity.Endpoint, 1)
	conn.Store.Find(&endpoints)
	r = response.GetAllEndpoints{Endpoints: endpoints}
	return r, nil
}
