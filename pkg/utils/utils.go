package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody reads the request body and unmarshals it into the provided interface
func ParseBody(r *http.Request, x interface{}) {
    // Read the entire request body
    body, err := io.ReadAll(r.Body)
    if err != nil {
        panic(err)
    }
    
    // Unmarshal JSON body into provided interface
    err = json.Unmarshal(body, x)
    if err != nil {
        panic(err)
    }
}