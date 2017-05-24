package action

import (
	"net/http"

	"go-get-it/infrastructure"
	"encoding/json"
)

func SaveAction(response http.ResponseWriter, request *http.Request){
	decoder := json.NewDecoder(req.Body)
	var t test_struct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	for key, value := range request.PostForm {
		infrastructure.Save(key, value[0])
	}

	response.Header().Set("Server", "Go-Get-It Server")
	response.WriteHeader(200)
	response.Write([]byte("Saved"))
}