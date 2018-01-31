package interactor

import (
	"github.com/sem-onyalo/application-dashboard/core/entity"
	"github.com/sem-onyalo/application-dashboard/service"
	"github.com/sem-onyalo/application-dashboard/service/request"
	"github.com/sem-onyalo/application-dashboard/service/response"
)

// Association is an interactor for associations
type Association struct {
	Database service.Database
}

// NewAssociation creates a reference to an association interactor
func NewAssociation(database service.Database) *Association {
	return &Association{Database: database}
}

// Create creates a new association
func (a Association) Create(request request.CreateAssociation) (response.CreateAssociation, error) {
	var r response.CreateAssociation

	conn, err := a.Database.NewConnection()
	if err != nil {
		return r, err
	}
	defer conn.Store.Close()

	var assocation = entity.Association{Name: request.Name}
	conn.Store.Create(&assocation)
	r.ID = assocation.ID
	return r, nil
}
