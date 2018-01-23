package service

import (
	"github.com/sem-onyalo/application-dashboard/service/response"
)

// Database is a boundary that allows you to interact with a database
type Database interface {
	NewConnection() (response.NewConnection, error)
}
