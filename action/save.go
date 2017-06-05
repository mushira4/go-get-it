package action

import (
	"net/http"
	"encoding/json"
	"log"

	"go-get-it/infrastructure"
)

type SaveRequest struct {
	keys map[string] interface{}
}

func SaveAction(response http.ResponseWriter, request *http.Request){
	decoder := json.NewDecoder(request.Body)
	var saveRequest SaveRequest
	err := decoder.Decode(&saveRequest.keys)
	if err != nil {
		panic(err)
	}
	defer request.Body.Close()

	for key, value := range saveRequest.keys {
		log.Printf("key[%s] value[%s]\n", key, value)
		infrastructure.Save(key, value.(string))
	}

	response.Header().Set("Server", "Go-Get-It Server")
	response.WriteHeader(200)
	response.Write([]byte("Saved"))
}