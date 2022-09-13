package configs

import (
	"os"
	"strings"
)

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

type Credentials struct {
	Dialect  string `json:"dialect"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"name"`
	Cluster  string `json:"cluster"`
}

func MakeURL(creds Credentials) (URL string) {

	var URLBuilder strings.Builder

	user := creds.User
	password := creds.Password
	host := creds.Host
	port := creds.Port
	name := creds.Name

	URLBuilder.WriteString("postgresql://" + user + ":" + password + "@" + host + ":" + port + "/" + name)

	return URLBuilder.String()
}

// GetConfig: method to fetch configuration
func GetConfig() *config { return &c }

// const: various consts converning the different results of the calls
const (
	STATUS_SUCCESS    string = "success"
	CODE_SUCCESS      int    = 200
	STATUS_EMPTY      string = "empty"
	CODE_EMPTY        int    = 204
	STATUS_BADREQUEST string = "failure"
	CODE_BADREQUEST   int    = 400
	STATUS_FAILURE    string = "failure"
	CODE_FAILURE      int    = 500
)
