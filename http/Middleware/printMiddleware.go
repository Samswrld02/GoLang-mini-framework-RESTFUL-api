package middleware

import (
	"log"
	"net/http"
)

func PrintMiddleware(w http.ResponseWriter, r *http.Request) {
	log.Printf("this works lol")
}
