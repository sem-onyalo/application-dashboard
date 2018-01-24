package interactor

import (
	"os"

	"github.com/sem-onyalo/application-dashboard/service/request"
	"github.com/sem-onyalo/application-dashboard/service/response"
)

// Config is an interactor for retrieving config values
type Config struct {
}

// NewConfig returns a reference to the config interactor
func NewConfig() *Config {
	return &Config{}
}

// GetValue retrieves a config value based on the specified request parameters
func (c Config) GetValue(request request.GetConfigValue) response.GetConfigValue {
	v := os.Getenv(request.Key)
	return response.GetConfigValue{Value: v}
}
