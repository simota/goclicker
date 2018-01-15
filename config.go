package goclicker

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	geoip2 "github.com/oschwald/geoip2-golang"
)

var database *gorm.DB
var geoDatabase *geoip2.Reader

func getConnectionString() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	protocol := "tcp"
	dbargs := "?charset=utf8&parseTime=true"

	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s", user, pass, protocol, host, port, dbname, dbargs)
}

// GetDatabase is
func GetDatabase() *gorm.DB {
	if database == nil {
		db, err := gorm.Open("mysql", getConnectionString())
		if err != nil {
			panic(err)
		}
		db.LogMode(true)
		database = db
	}
	return database
}

// GetGeoDatabase is
func GetGeoDatabase() *geoip2.Reader {
	if geoDatabase == nil {
		db, err := geoip2.Open("GeoLite2-Country.mmdb")
		if err != nil {
			log.Fatal(err)
		}
		geoDatabase = db
	}
	return geoDatabase
}

// GetAPIKey is
func GetAPIKey() string {
	return os.Getenv("API_KEY")
}

// GetEndpoint is
func GetEndpoint() string {
	return os.Getenv("ENDPOINT")
}
