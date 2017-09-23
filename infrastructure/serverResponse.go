package infrastructure

import "net/http"

func writeDefaultHeader(response http.ResponseWriter) {
	response.Header().Set("Server", "Go-Get-It Server")
}

func WriteOK(response http.ResponseWriter) {
	writeDefaultHeader(response)
	response.WriteHeader(200)
}
