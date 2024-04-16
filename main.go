package main

import (
	"dsr-mail-service/config"
	"dsr-mail-service/endpoints"
	"fmt"
	"net/http"
	"os"
)

// func handleRequests() {
// 	http.Handle("/creditscore", http.HandlerFunc(getCreditScore))
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

func main() {
	config.LoadConfig()
	// mailing.SendEmail()
	endpoints.HandleRequests()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
