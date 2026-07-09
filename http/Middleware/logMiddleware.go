package middleware

import (
	"log"
	"net/http"
)

// take route to handle middleware
func LogMiddleware(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v", r.Method, r.URL.Path)
}
