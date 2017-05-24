package main

import (
	//Go libs
	"log"
	"net/http"

	//Internal
	"go-get-it/route"
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

	route.Routes()

	http.ListenAndServe(":" + Config.Port, nil)
}
