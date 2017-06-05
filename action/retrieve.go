package action

import (
	"log"
	"net/http"
	"strings"

	"go-get-it/infrastructure"
)

func RetrieveAction(response http.ResponseWriter, request *http.Request){
	query := request.URL.Query().Get("query")

	var searchQuery string

	if len(strings.TrimSpace(query)) == 0 {
		searchQuery = "*"
	} else {
		searchQuery = "*" + searchQuery + "*"
	}

	values := infrastructure.Retrieve(query)
        for value := range values {
		log.Println(value + " - " + values[value])
	}

	infrastructure.WriteOK(response)
	response.Write([]byte("Retrieved"))
}
