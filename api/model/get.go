
package model

import (
	"encoding/json"
	"net/http"
	"github.com/VladRomanciuc/Go-classes/api/views"
)

func Get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
    if r.Method == http.MethodGet {
		data := views.Response{
		Code: http.StatusOK,
		Body: "",
		}
		json.NewEncoder(w).Encode(data)
	}
}