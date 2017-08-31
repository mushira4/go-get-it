package action

import (
	"net/http"
	"encoding/json"
	"log"
	"fmt"

	"go-get-it/infrastructure"
)

type SaveRequest struct {
	keys map[string] interface{}
}

/**
 * SaveAction is the action endpoint responsible for save data into the storage.
 */
func SaveAction(response http.ResponseWriter, request *http.Request){
	decoder := json.NewDecoder(request.Body)
	var saveRequest SaveRequest
	err := decoder.Decode(&saveRequest.keys)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer request.Body.Close()

	for key, value := range saveRequest.keys {
		log.Printf("key[%s] value[%s]\n", key, value)
		infrastructure.Save(key, value.(string))
	}

	infrastructure.WriteOK(response)
	response.Write([]byte("Saved"))
}