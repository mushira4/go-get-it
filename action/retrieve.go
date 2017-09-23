package action

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"go-get-it/infrastructure"
)

/**
 * RetrieveAction is the action endpoint responsible for retrieve the stored valid data.
 */
func RetrieveAction(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query().Get("query")

	var searchQuery string

	if len(strings.TrimSpace(query)) == 0 {
		searchQuery = "*"
	} else {
		searchQuery = "*" + query + "*"
	}

	searchRetrieve := infrastructure.Retrieve(searchQuery)
	for value := range searchRetrieve {
		log.Println(value + " - " + searchRetrieve[value])
	}

	infrastructure.WriteOK(response)

	values, err := json.Marshal(searchRetrieve)
	if err != nil {
		fmt.Println(err)
		return
	}

	response.Write([]byte(string(values)))
}
