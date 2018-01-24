package service

import (
	"github.com/sem-onyalo/application-dashboard/service/request"
	"github.com/sem-onyalo/application-dashboard/service/response"
)

// Config is a boundary that allows you to retrieve configuration values
type Config interface {
	GetValue(request request.GetConfigValue) response.GetConfigValue
}
