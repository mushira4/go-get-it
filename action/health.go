package action

import (
	"net/http"

	"go-get-it/infrastructure"
)

/**
 * HealthAction is the action endpoint responsible for confirm the server health.
 */
func HealthAction(response http.ResponseWriter, request *http.Request){
	infrastructure.WriteOK(response)
	response.Write([]byte("OK"))
}
