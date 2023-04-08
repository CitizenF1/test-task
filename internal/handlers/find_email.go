package handlers

import (
	"fmt"
	"io"
	"net/http"
	"rest-api/internal/service"
)

func HandleFindEmail(w http.ResponseWriter, r *http.Request) {
	// поиск email адресов в теле запроса
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	emails := service.FindEmail(string(body))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", emails)))
}

func HandleFindIIN(w http.ResponseWriter, r *http.Request) {
	// поиск ИИН в теле запроса
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	iins := service.Findiin(string(body))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", iins)))
}
