package goclicker

import (
	"testing"

	"github.com/simota/goclicker"
)

func TestGetCountry(t *testing.T) {
	var country string
	country = goclicker.GetCountry("122.212.183.146")
	if country != "JP" {
		t.Errorf("Error TestGetCountry JP")
	}
	country = goclicker.GetCountry("122.212.183.146:49684")
	if country != "JP" {
		t.Errorf("Error TestGetCountry JP")
	}
}

func TestGetRefererHost(t *testing.T) {
	var host string
	host = goclicker.GetRefererHost("https://www.google.co.jp/search?q=wefwe&ie=utf-8&oe=utf-8&client=firefox-b-ab&gfe_rd=cr&ei=xhE9WPGqCIXf8AfSr5DACA")
	if host != "www.google.co.jp" {
		t.Errorf("Error TestGetRefererHost www.google.co.jp")
	}
	host = goclicker.GetRefererHost("")
	if host != "" {
		t.Errorf("Error TestGetRefererHost Unknown")
	}
}
