package routes

import (
	"encoding/json"
	"net/http"
)

func Regrouter() *http.ServeMux {
	mux := &http.ServeMux{}

	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Канадская береговая")
	})

	return mux
}
