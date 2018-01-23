package interactor

import (
	"fmt"

	"github.com/sem-onyalo/application-dashboard/service/response"

	"github.com/jinzhu/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "d1e4211d2a825210a6579d18c5a01fd25494ed42"
	dbname   = "appdash"
)

// Database is an interactor for interacting with the Database entity
type Database struct {
	ConnStr string
}

// NewDatabase creates a new database object instance with connection string info
func NewDatabase() (*Database, error) {
	// TODO: pull from config
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	return &Database{ConnStr: psqlInfo}, nil
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
