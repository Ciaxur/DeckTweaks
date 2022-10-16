package ping

import "github.com/gorilla/mux"

func CreateRoute(r *mux.Router) error {
	r.HandleFunc("", handlePing).Methods("GET")
	return nil
}
