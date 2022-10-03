package main

import (
	"dddx-lp-data/initializers"
	"dddx-lp-data/routers"
	"log"
	"net/http"
	"os"
)

// init is a function automatically run before main func.
// init contains 02 actions including: (1) Loading .env file to use;
// and (2) Connecting to Ethereum client in public BSC RPC node.
func init() {
	initializers.LoadEnv()
	initializers.ConnectToEthClient()
}

func main() {
	//
	r := routers.Routers()

	err := http.ListenAndServe(os.Getenv("PORT"), r)
	if err != nil {
		// log.Fatal("Failed to start server")
		log.Fatal(err)
	}
}
