package main

import (
	//Go libs
	"log"
	"net/http"
	"strconv"

	//Internal
	"go-get-it/action"
	"go-get-it/config"
)

func init() {
	log.SetPrefix("INFO:")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	log.Println("############################")
	log.Println("Initializing System")
	log.Println("############################")
}

func main() {
	log.Println("Finally Started")

	Config, _ := config.ReadConfig()

	http.HandleFunc("/health", action.HealthAction)
	http.HandleFunc("/send", action.SendAction)
	http.HandleFunc("/retrieve", action.RetrieveAction)
	http.ListenAndServe(":"+strconv.Itoa(Config.Port), nil)
}
