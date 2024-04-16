package endpoints

import (
	"net/http"
)

func GetHome() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		GetHandler(w, r)
		HandleJSONResponse(w, http.StatusOK, map[string]string{"message": "Welcome to the home page"})
	})
}
