package handlers

import (
	"fmt"
	"net/http"
	"rest-api/internal/service"

	"github.com/gorilla/mux"
)

func HandleFindSelf(w http.ResponseWriter, r *http.Request) {
	substr := mux.Vars(r)["substr"]

	matches, err := service.FindSelf(substr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", matches)))
}
