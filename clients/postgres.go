package clients

import (
	cfg "sybo/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	l "gorm.io/gorm/logger"
)

var db *gorm.DB

func GetCon() (*gorm.DB, error) {

	if db == nil {

		dns := cfg.MakeURL(cfg.GetConfig().DB)

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
