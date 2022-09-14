package clients

import (
	"strings"
	cfg "sybo/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	l "gorm.io/gorm/logger"
)

// Keep the gorm connection
var db *gorm.DB

// makePostgresConnectionString: convenience function to build the connection string. Could be different for other db types.
func makePostgresConnectionString(creds cfg.Credentials) (URL string) {

	var URLBuilder strings.Builder

	user := creds.User
	password := creds.Password
	host := creds.Host
	port := creds.Port
	name := creds.Name

	URLBuilder.WriteString("postgresql://" + user + ":" + password + "@" + host + ":" + port + "/" + name)

	return URLBuilder.String()
}

// GetCon: Setup the connection or return the existing one.
func GetCon() (*gorm.DB, error) {

	if db == nil {

		//Get the connection string
		dns := makePostgresConnectionString(cfg.GetConfig().DB)

		// Open the conneciton to the given DB
		con, err := gorm.Open(postgres.Open(dns), &gorm.Config{
			Logger: l.Default.LogMode(l.Silent),
		})
		if err != nil {
			return db, err
		}

		db = con

		return db, nil
	}

	return db, nil
}
