package utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func GetAllFarmingPoolLength() int {
	poolLength, err := DDDXVoter.Length(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	return int(poolLength.Uint64())
}
