package interactor

import (
	"github.com/sem-onyalo/application-dashboard/service"
	"github.com/sem-onyalo/application-dashboard/service/request"
	"github.com/sem-onyalo/application-dashboard/service/response"

	"github.com/jinzhu/gorm"
)

const (
	databaseURLConfigKey = "APPDASH_DATABASE_URL"
)

// Database is an interactor for interacting with a database
type Database struct {
	ConnStr string
}

// NewDatabase returns a reference to the database interactor with the connection string info set
func NewDatabase(config service.Config) (*Database, error) {
	dbURLConfig := config.GetValue(request.GetConfigValue{Key: databaseURLConfigKey})
	return &Database{ConnStr: dbURLConfig.Value}, nil
}

// NewConnection creates a new connection to the database entity. Caller is responsible for closing connection
func (d Database) NewConnection() (response.NewConnection, error) {
	db, err := gorm.Open("postgres", d.ConnStr)

	if err != nil {
		return response.NewConnection{}, err
	}

	db.SingularTable(true)
	return response.NewConnection{Store: db}, nil
}
