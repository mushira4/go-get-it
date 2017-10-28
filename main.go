package main

import (
	"net/http"
	"runtime"

	"go-get-it/controller"
	"go-get-it/infrastructure/config"
)

func main() {
	/**
	 * Setting up the max number of possible cores to use.
	 */
	runtime.GOMAXPROCS(runtime.NumCPU())

	controller.Routes()

	http.ListenAndServe(":"+config.Config.Port, nil)
}
