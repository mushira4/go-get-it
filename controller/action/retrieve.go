package action

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"go-get-it/infrastructure"
	"go-get-it/infrastructure/logger"
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
	actionLogger := logger.SimpleLogger()
	for value := range searchRetrieve {
		actionLogger.Append(value, searchRetrieve[value])
	}
	actionLogger.Info()

	infrastructure.WriteOK(response)

	values, err := json.Marshal(searchRetrieve)
	if err != nil {
		fmt.Println(err)
		return
	}

	response.Write([]byte(string(values)))
}
