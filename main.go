package main

import (
	"dddx-lp-data/initializers"
	"dddx-lp-data/routers"
	"log"
	"net/http"
	"os"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToEthClient()
}

func main() {
	r := routers.Routers()

	err := http.ListenAndServe(os.Getenv("PORT"), r)
	if err != nil {
		// log.Fatal("Failed to start server")
		log.Fatal(err)
	}
}
