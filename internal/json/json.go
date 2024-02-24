package json

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func Decode(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
