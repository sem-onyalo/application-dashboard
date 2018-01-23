package response

import (
	"github.com/sem-onyalo/application-dashboard/core/entity"
)

// GetAllEndpoints represents the response to a GetAllEndpoints request
type GetAllEndpoints struct {
	Endpoints []entity.Endpoint
}
