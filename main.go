package main

import (
  //Go libs
  "io/ioutil"
  "log"
  "net/http"
  "strconv"

  //External libs
  "github.com/ghodss/yaml"
)

type AppConfig struct{
  Port int 
  Log string
}

func init(){
  log.SetPrefix("INFO:")
  log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
  log.Println("############################")
  log.Println("Initializing System")
  log.Println("############################")
}


func main(){
  log.Println("Finally Started")

  data, err := ioutil.ReadFile("/home/ekashivagui/workspace/src/go-get-it/config/local.yaml")
  if err != nil {
    panic(err)
  }

  var config AppConfig
  err = yaml.Unmarshal(data, &config)
  log.Println("############################")
  log.Println("Server Config")
  log.Println("Port: " + strconv.Itoa(config.Port))
  log.Println("Log Path: " + config.Log)
  log.Println("############################")

  http.HandleFunc("/health", healthFunction)
  http.ListenAndServe(":8080", nil)
}

func healthFunction(response http.ResponseWriter, request *http.Request){
  response.Header().Set("Server", "Go-Get-It Server")
  response.WriteHeader(200)
  response.Write([]byte("OK"))
}

