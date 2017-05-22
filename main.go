package main

import (
	//Go libs
	"log"
	"net/http"

	//Internal
	"go-get-it/action"
	"go-get-it/config"
)

var Config *config.AppConfig
var ConfigError error

func init() {
	log.SetPrefix("INFO:")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	log.Println("############################")
	log.Println("Initializing System")
	log.Println("############################")

        Config, ConfigError = config.ReadConfig()
        if ConfigError != nil {
          panic(ConfigError)
        }
}

func main() {
	log.Println("Finally Started")
	http.HandleFunc("/health", action.HealthAction)
	http.HandleFunc("/save", action.SaveAction)
	http.HandleFunc("/retrieve", action.RetrieveAction)
	http.ListenAndServe(":" + Config.Port, nil)
}
