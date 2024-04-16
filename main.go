package main

import (
	"dsr-mail-service/config"
	"dsr-mail-service/endpoints"
	"net/http"
)

// func handleRequests() {
// 	http.Handle("/creditscore", http.HandlerFunc(getCreditScore))
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

func main() {
	config.LoadConfig()
	// mailing.SendEmail()
	endpoints.HandleRequests()

	http.ListenAndServe(":8080", nil)
}
