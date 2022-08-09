package corsutils

import (
	"net/http"
	"regexp"

	"github.com/spf13/viper"
)

func AllowOrigin(origin string) bool {
	allowedOrigin := viper.GetString("allowed_origin")
	if allowedOrigin == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(allowedOrigin, origin); matched {
		return true
	}
	return false
}

// Cors middleware allow requests coming from origin which allowOrigin func returns true for it
func Cors(h http.Handler, allowOrigin func(string) bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}
