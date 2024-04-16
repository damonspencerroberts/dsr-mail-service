package endpoints

import (
	"dsr-mail-service/mailing"
	"dsr-mail-service/types"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func PostEmailRequest() {
	http.HandleFunc("/send_email", func(w http.ResponseWriter, r *http.Request) {
		err := validateToken(r)
		if err != nil {
			errorJSON := []byte(fmt.Sprintf(`{"error": {"message": "%s"}}`, err.Error()))
			http.Error(w, string(errorJSON), http.StatusUnauthorized)
			return
		}
		var requestBody types.EmailRequest
		if err := decodeAndValidateRequest(r, &requestBody); err != nil {
			errorJSON := []byte(fmt.Sprintf(`{"error": {"message": "%s"}}`, err.Error()))
			http.Error(w, string(errorJSON), http.StatusBadRequest)
			return
		}

		PostHandler(w, r)

		mailing.SendEmail(requestBody)
		HandleJSONResponse(w, http.StatusOK, requestBody)
	})
}
