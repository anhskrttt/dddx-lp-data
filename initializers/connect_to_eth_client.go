package initializers

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var Client *ethclient.Client

func ConnectToEthClient() {
	var err error
	Client, err = ethclient.Dial("https://bsc-dataseed4.ninicoin.io")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("We have a connection")
}
