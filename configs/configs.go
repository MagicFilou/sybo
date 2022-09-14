package configs

import (
	"os"
)

// Holds the config object
var c config

// init: setup the configuration to get data from the env variables
func init() {
	c.ENV = os.Getenv("ENV")
	c.Port = os.Getenv("PORT")

	c.DB.User = os.Getenv("DB_USER")
	c.DB.Password = os.Getenv("DB_PASSWORD")
	c.DB.Host = os.Getenv("DB_HOST")
	c.DB.Port = os.Getenv("DB_PORT")
	c.DB.Name = os.Getenv("DB_NAME")
}

// config: srtuct to hold configuration
type config struct {
	ENV  string
	Port string
	DB   Credentials
}

// Credentials: struct to hold db credentials
type Credentials struct {
	Dialect  string `json:"dialect"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"name"`
	Cluster  string `json:"cluster"`
}

// GetConfig: function to fetch configuration
func GetConfig() *config { return &c }

// const: various consts for the service
const (
	SEPARATOR string = ","
)
