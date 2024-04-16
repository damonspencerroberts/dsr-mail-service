package endpoints

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type EmailRequest struct {
	ToEmail   string `json:"to_email" validate:"required,email"`
	Subject   string `json:"subject" validate:"required"`
	Message   string `json:"message" validate:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

var validate = validator.New()

func PostEmailRequest() {
	http.HandleFunc("/send_email", func(w http.ResponseWriter, r *http.Request) {
		var requestBody EmailRequest
		if err := decodeAndValidateRequest(r, &requestBody); err != nil {
			errorJSON := []byte(fmt.Sprintf(`{"error": {"message": "%s"}}`, err.Error()))
			http.Error(w, string(errorJSON), http.StatusBadRequest)
			return
		}

		body, err := PostHandler(w, r)

		if err != nil {
			return
		}

		HandleJSONResponse(w, http.StatusOK, map[string]string{"body": string(body)})
	})
}
