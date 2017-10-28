package action

import (
	"encoding/json"
	"net/http"

	"go-get-it/infrastructure/logger"
	"go-get-it/infrastructure/storage"
)

type SaveRequest struct {
	keys map[string]interface{}
}

/**
 * SaveAction is the action endpoint responsible for save data into the storage.
 */
func SaveAction(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var saveRequest SaveRequest
	err := decoder.Decode(&saveRequest.keys)
	if err != nil {
		logger.EmptyLogger().Panic(err)
	}
	defer request.Body.Close()

	for key, value := range saveRequest.keys {
		logger.Append(key, value.(string)).Info()
		infrastructure.Save(key, value.(string))
	}

	infrastructure.WriteOK(response)
	response.Write([]byte("Saved"))
}
