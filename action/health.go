package action

import "net/http"

func HealthAction(response http.ResponseWriter, request *http.Request){
	response.Header().Set("Server", "Go-Get-It Server")
	response.WriteHeader(200)
	response.Write([]byte("OK"))
}
