package service

import (
	"github.com/sem-onyalo/application-dashboard/service/request"
	"github.com/sem-onyalo/application-dashboard/service/response"
)

// Resolution is a boundary to interacting with resolutions
type Resolution interface {
	Create(request.CreateResolution) (response.CreateResolution, error)
}
