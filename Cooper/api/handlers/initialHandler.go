package handlers

import (
	"net/http"

	opc "example.com/coomper/opc"
)

func Initial(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	_, err := opc.InitServer([]byte(r.FormValue("data")))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}
