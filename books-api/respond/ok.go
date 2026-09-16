package respond

import (
	"encoding/json"
	"net/http"
)

func OK(w http.ResponseWriter, data any) {
	if data == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return
	}
}
