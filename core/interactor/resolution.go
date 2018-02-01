package interactor

import (
	"errors"

	"github.com/sem-onyalo/application-dashboard/core/entity"
	"github.com/sem-onyalo/application-dashboard/service"
	"github.com/sem-onyalo/application-dashboard/service/request"
	"github.com/sem-onyalo/application-dashboard/service/response"
)

// Resolution is an interactor for interating with resolutions
type Resolution struct {
	Database service.Database
}

// NewResolution creates a reference to the resolution interactor
func NewResolution(database service.Database) *Resolution {
	return &Resolution{Database: database}
}

// Create creates a resolution
func (r Resolution) Create(req request.CreateResolution) (response.CreateResolution, error) {
	var res response.CreateResolution

	conn, err := r.Database.NewConnection()
	if err != nil {
		return res, err
	}
	defer conn.Store.Close()

	var resolution = entity.Resolution{
		Name:    req.Name,
		Details: req.Details,
	}
	conn.Store.Create(&resolution)
	if resolution.ID <= 0 {
		return res, errors.New("Create resolution failed")
	}

	res.ID = resolution.ID
	return res, nil
}
