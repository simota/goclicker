package goclicker

import (
	"strings"
	"time"
)

// Link is
type Link struct {
	ID        uint64    `gorm:"primary_key" json:"linkId"`
	Key       string    `gorm:"type:varchar(100);not null;unique_index" json:"nil"`
	URL       string    `gorm:"type:varchar(255);not null;index" json:"longUrl"`
	Endpoint  string    `gorm:"-" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// CreateEndpoint is
func (e *Link) CreateEndpoint(s string) {
	e.Endpoint = s + "/" + e.Key
}

// AccessLog is
type AccessLog struct {
	ID          uint64    `gorm:"primary_key" json:"nil"`
	LinkID      uint64    `gorm:"index;not null" json:"nil"`
	UserAgent   string    `gorm:"type:varchar(255)" json:"userAgent"`
	Referer     string    `gorm:"type:varchar(255)" json:"referer"`
	RefererHost string    `gorm:"type:varchar(255)" json:"refererHost"`
	IPAddress   string    `gorm:"type:varchar(30)" json:"ipaddress"`
	Country     string    `gorm:"type:varchar(100)" json:"country"`
	Platform    string    `gorm:"type:varchar(100)" json:"platform"`
	Browser     string    `gorm:"type:varchar(100)" json:"browser"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// LinkResponse is
type LinkResponse struct {
	Link
	Logs []AccessLog `json:"logs"`
}

// URLRequest is
type URLRequest struct {
	LongURL string `json:"longUrl"`
}

// UserAgent is
type UserAgent struct {
	ua string
}

// NewUserAgent is
func NewUserAgent(s string) UserAgent {
	ua := UserAgent{}
	ua.ua = s
	return ua
}

// String is
func (u *UserAgent) String() string {
	return u.ua
}

// Platform is
func (u *UserAgent) Platform() string {
	if strings.Contains(u.ua, "iPod touch") {
		return "iPod touch"
	}
	if strings.Contains(u.ua, "iPad") {
		return "iPad"
	}
	if strings.Contains(u.ua, "iPhone") {
		return "iPhone"
	}
	if strings.Contains(u.ua, "Android") {
		return "Android"
	}
	if strings.Contains(u.ua, "Windows") {
		return "Windows"
	}
	if strings.Contains(u.ua, "Macintosh") {
		return "Macintosh"
	}
	if strings.Contains(u.ua, "Linux") {
		return "Linux"
	}
	return "Unknown"
}

// GetBrowser is
func (u *UserAgent) Browser() string {
	if strings.Contains(u.ua, "Edge") {
		return "Edge"
	}
	if strings.Contains(u.ua, "Trident") {
		return "IE"
	}
	if strings.Contains(u.ua, "Firefox") {
		return "Firefox"
	}
	if strings.Contains(u.ua, "Chrome") {
		return "Chrome"
	}
	if strings.Contains(u.ua, "Safari") {
		return "Safari"
	}
	return "Unknown"
}
