package service

import "github.com/sem-onyalo/application-dashboard/service/request"

// App is a boundary to web application functionality
type App interface {
	Start(request request.StartApp)
}
