package action

import (
	"net/http"

	"go-get-it/infrastructure"
	"log"
)

func RetrieveAction(response http.ResponseWriter, request *http.Request){
	values := infrastructure.Retrieve("*")
        for value := range values {
		log.Println(value)
	}

	response.Header().Set("Server", "Go-Get-It Server")
	response.WriteHeader(200)
	response.Write([]byte("Retrieved"))
}
