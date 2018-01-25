package service

import (
	"github.com/sem-onyalo/application-dashboard/service/request"
	"github.com/sem-onyalo/application-dashboard/service/response"
)

// App is a boundary to web application functionality
type App interface {
	Start(request request.StartApp) response.StartApp
}
