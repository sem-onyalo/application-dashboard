package service

import (
	"github.com/sem-onyalo/application-dashboard/service/request"
	"github.com/sem-onyalo/application-dashboard/service/response"
)

// Association is a boundary to interacting with an association
type Association interface {
	Create(request.CreateAssociation) (response.CreateAssociation, error)
}
