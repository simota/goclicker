package goclicker

import (
	"testing"

	"github.com/simota/goclicker"
)

func TestGetPlatform(t *testing.T) {
	var ua goclicker.UserAgent

	ua = goclicker.NewUserAgent("Mozilla/5.0 (iPhone; CPU iPhone OS 10_0_2 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) Version/10.0 Mobile/14A456 Safari/602.1")
	if ua.Platform() != "iPhone" {
		t.Errorf("Error TestGetPlatform iPhone")
	}

	ua = goclicker.NewUserAgent("Mozilla/5.0 (iPod touch; CPU iPhone OS 10_0_2 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) Version/10.0 Mobile/14A456 Safari/602.1")
	if ua.Platform() != "iPod touch" {
		t.Errorf("Error TestGetPlatform iPod touch")
	}

	ua = goclicker.NewUserAgent("Mozilla/5.0 (Linux; Android 7.0; Nexus 5X Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.124 Mobile Safari/537.36")
	if ua.Platform() != "Android" {
		t.Errorf("Error TestGetPlatform Android")
	}

	ua = goclicker.NewUserAgent("Mozilla/5.0 (iPad; CPU OS 10_0_2 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) Version/10.0 Mobile/14A456 Safari/602.1")
	if ua.Platform() != "iPad" {
		t.Errorf("Error TestGetPlatform iPad")
	}

	ua = goclicker.NewUserAgent("Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36")
	if ua.Platform() != "Windows" {
		t.Errorf("Error TestGetPlatform Windows")
	}

	ua = goclicker.NewUserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36")
	if ua.Platform() != "Macintosh" {
		t.Errorf("Error TestGetPlatform Macintosh")
	}

	ua = goclicker.NewUserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36")
	if ua.Platform() != "Linux" {
		t.Errorf("Error TestGetPlatform Linux")
	}

	ua = goclicker.NewUserAgent("Unknown")
	if ua.Platform() != "Unknown" {
		t.Errorf("Error TestGetPlatform Unknown")
	}
}

func TestGetBrowser(t *testing.T) {
	var ua goclicker.UserAgent

	ua = goclicker.NewUserAgent("Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.79 Safari/537.36 Edge/14.14393")
	if ua.Browser() != "Edge" {
		t.Errorf("Error TestGetBrowser Edge")
	}

	ua = goclicker.NewUserAgent("Mozilla/5.0 (Windows NT 10.0; Trident/7.0; rv:11.0) like Gecko")
	if ua.Browser() != "IE" {
		t.Errorf("Error TestGetBrowser IE")
	}

	ua = goclicker.NewUserAgent("Mozilla/5.0 (Windows NT 10.0; rv:49.0) Gecko/20100101 Firefox/49.0")
	if ua.Browser() != "Firefox" {
		t.Errorf("Error TestGetBrowser Firefox")
	}

	ua = goclicker.NewUserAgent("Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36")
	if ua.Browser() != "Chrome" {
		t.Errorf("Error TestGetBrowser Chrome")
	}

	ua = goclicker.NewUserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12) AppleWebKit/602.1.50 (KHTML, like Gecko) Version/10.0 Safari/602.1.50")
	if ua.Browser() != "Safari" {
		t.Errorf("Error TestGetBrowser Safari")
	}
}
