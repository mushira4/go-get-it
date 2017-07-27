package action

import (
	"net/http"

	"go-get-it/infrastructure"
)

func HealthAction(response http.ResponseWriter, request *http.Request){
	infrastructure.WriteOK(response)
	response.Write([]byte("OK"))
}
