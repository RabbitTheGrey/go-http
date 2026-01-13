package internal

import (
	"encoding/json"
	"net/http"
)

// Make application/json response with encoded body
// and set http status code
func JsonResponse(body any, status int, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}
