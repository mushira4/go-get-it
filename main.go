package main

import (
  "log"
  "net/http"
)

func init(){
  log.SetPrefix("INFO:")
  log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
  log.Println("############################")
  log.Println("Initializing System")
  log.Println("############################")
}


func main(){
  log.Println("Finally Started") 
  http.HandleFunc("/health", healthFunction)
  http.ListenAndServe(":8080", nil)
}

func healthFunction(response http.ResponseWriter, request *http.Request){
  response.Header().Set("Server", "Go-Get-It Server")
  response.WriteHeader(200)
  response.Write([]byte("OK"))
}

