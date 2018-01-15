package goclicker

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net"
	"net/url"
	"strings"

	uuid "github.com/satori/go.uuid"
)

// CreateKey is
func CreateKey(value string) string {
	s := sha256.New()
	u4 := uuid.NewV4()
	io.WriteString(s, value+u4.String())
	result := fmt.Sprintf("%x", s.Sum(nil))
	return result[0:7]
}

// GetCountry is
func GetCountry(ipaddr string) string {
	db := GetGeoDatabase()
	ip := net.ParseIP(strings.Split(ipaddr, ":")[0])
	record, err := db.City(ip)
	if err != nil {
		return "Unknown"
	}
	return record.Country.IsoCode
}

// GetRefererHost is
func GetRefererHost(s string) string {
	u, err := url.Parse(s)
	if err != nil {
		return ""
	}
	return u.Host
}
