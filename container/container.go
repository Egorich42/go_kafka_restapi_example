package container

import (
	"github.com/Egorich42/testserver/config"
)

// Container represents an interface for accessing the data which sharing in overall application
type Container interface {
	GetConfig() *config.AppConfig
}

// container struct is for sharing data such as the setting of application and logger in overall this application
type container struct {
	config *config.AppConfig
}

// NewContainer is constructor
func NewContainer(config *config.AppConfig) Container {
	return &container{config: config}
}

// GetConfig returns the object of configuration
func (c *container) GetConfig() *config.AppConfig {
	return c.config
}
