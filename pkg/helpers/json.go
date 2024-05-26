package helpers

import (
	"encoding/json"
	"net/http"
)

func DecodeJSONRequest(r *http.Request, target interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	return decoder.Decode(target)
}
