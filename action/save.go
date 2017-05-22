package action

import (
	"net/http"

	"go-get-it/infrastructure"
)

func SaveAction(response http.ResponseWriter, request *http.Request){
	request.PostFormValue("Mandrake Call")

	for key, value := range request.PostForm {
		infrastructure.Save(key, value[0])
	}

	response.Header().Set("Server", "Go-Get-It Server")
	response.WriteHeader(200)
	response.Write([]byte("Saved"))
}