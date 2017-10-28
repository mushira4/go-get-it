package controller

import (
	"net/http"

	"go-get-it/controller/action"
)

/**
 * Routes define all the routes used by the server
 */
func Routes() {
	http.HandleFunc("/health", action.HealthAction)
	http.HandleFunc("/save", action.SaveAction)
	http.HandleFunc("/retrieve", action.RetrieveAction)
}
