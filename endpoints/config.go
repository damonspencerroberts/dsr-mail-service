package endpoints

import (
	"fmt"
	"io"
	"net/http"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	fmt.Println("Received POST request:")
	fmt.Println(string(body))

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "POST request received successfully")
}
