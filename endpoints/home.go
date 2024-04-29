package endpoints

import (
	"net/http"
)

func getHomeRequest(w http.ResponseWriter, r *http.Request) {
	GetHandler(w, r)
	HandleJSONResponse(w, http.StatusOK, map[string]string{"message": "Welcome to the home page"})
}

func GetHome() {
	http.HandleFunc("/", EnableCors(getHomeRequest))
}
