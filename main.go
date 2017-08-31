package main

import (
	//Go libs
	"log"
	"net/http"
	"runtime"

	//Internal
	"go-get-it/config"
	"go-get-it/route"
)

var Config *config.AppConfig

func init() {
	log.SetPrefix("INFO:")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	log.Println("############################")
	log.Println("Initializing System")
	log.Println("############################")
}

func main() {
	log.Println("############################")
	log.Println("Finally Started")
	log.Println("############################")

	/**
	 * Setting up the max number of possible cores to use.
	 */
	runtime.GOMAXPROCS(runtime.NumCPU())

	route.Routes()

	http.ListenAndServe(":"+Config.Port, nil)
}
