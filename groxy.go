package groxy

import "github.com/sirupsen/logrus"

type Config struct {
	Development bool

	// AirTable configuration
	AppID  string
	APIKey string

	Host string
	Port int

	Logger *logrus.Logger
}
