package handlers

import (
	"fmt"
	"io"
	"net/http"
	"rest-api/internal/service"
)

func HandleFindSubStr(w http.ResponseWriter, r *http.Request) {
	// поиск подстроки в теле запроса
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	subString := service.FindMaxSubstring(string(body))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", subString)))
}
