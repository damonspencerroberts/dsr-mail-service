package endpoints

import (
	"net/http"
)

func PostEmailRequest() {
	http.HandleFunc("/send_email", func(w http.ResponseWriter, r *http.Request) {
		body, err := PostHandler(w, r)
		if err != nil {
			return
		}
		HandleJSONResponse(w, http.StatusOK, map[string]string{"body": string(body)})
	})
}
