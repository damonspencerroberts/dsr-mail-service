package endpoints

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HandleJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)
	w.Write(jsonData)
}

func GetHandler(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return nil, fmt.Errorf("method not allowed")
	}

	return nil, nil
}

func PostHandler(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return nil, fmt.Errorf("method not allowed")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return nil, err
	}

	return body, nil
}
