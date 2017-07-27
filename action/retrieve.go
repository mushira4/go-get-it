package action

import (
	"log"
	"net/http"
	"strings"
	"encoding/json"
	"fmt"

	"go-get-it/infrastructure"

)


func RetrieveAction(response http.ResponseWriter, request *http.Request){
	query := request.URL.Query().Get("query")

	var searchQuery string

	if len(strings.TrimSpace(query)) == 0 {
		searchQuery = "*"
	} else {
		searchQuery = "*" + query + "*"
	}

	values := infrastructure.Retrieve(searchQuery)
        for value := range values {
		log.Println(value + " - " + values[value])
	}

	infrastructure.WriteOK(response)

	b, err := json.Marshal(values)
	if err != nil {
		fmt.Println(err)
		return
	}

	response.Write([]byte(string(b)))
}
