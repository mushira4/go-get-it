package action

import (
	"log"
	"net/http"

	"go-get-it/infrastructure"
)

func RetrieveAction(response http.ResponseWriter, request *http.Request){
	query := request.URL.Query().Get("query")
	values := infrastructure.Retrieve(query)
        for value := range values {
		log.Println(value + " - " + values[value])
	}

	response.Header().Set("Server", "Go-Get-It Server")
	response.WriteHeader(200)
	response.Write([]byte("Retrieved"))
}
