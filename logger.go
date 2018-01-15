package goclicker

import "net/http"

// Logging is
func Logging(id uint64, r *http.Request) {
	referer := r.Header.Get("Referer")
	ua := NewUserAgent(r.Header.Get("User-Agent"))
	go func() {
		repo := NewAccessLogRepository()
		log := AccessLog{}
		log.LinkID = id
		log.UserAgent = ua.String()
		log.Referer = referer
		log.RefererHost = GetRefererHost(referer)
		log.IPAddress = r.RemoteAddr

		log.Platform = ua.Platform()
		log.Browser = ua.Browser()
		log.Country = GetCountry(r.RemoteAddr)

		repo.Save(log)
	}()
}
