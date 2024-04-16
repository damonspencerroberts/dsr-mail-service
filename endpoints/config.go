package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
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

	return nil, nil
}

func decodeAndValidateRequest(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return fmt.Errorf("error decoding request body: %s", err)
	}

	if err := validate.Struct(v); err != nil {
		return fmt.Errorf("invalid request body: %s", err)
	}

	return nil
}

func validateToken(r *http.Request) error {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return fmt.Errorf("authorization header is missing")
	}
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return fmt.Errorf("invalid authorization format")
	}

	token := parts[1]

	if !isValidToken(token) {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func isValidToken(token string) bool {
	return token == os.Getenv("API_TOKEN")
}
