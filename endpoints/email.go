package endpoints

import (
	"net/http"
)

func PostEmailRequest() {
	http.HandleFunc("/post", PostHandler)
}
