package utils

import (
	"dddx-lp-data/abigen/voter"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func GetAllFarmingPoolLength(voterInstance *voter.Voter) int {
	poolLength, err := voterInstance.Length(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	return int(poolLength.Uint64())
}

func GetAllGaugesOf(voter *voter.Voter) (int, []string) {
	var pools []string
	var gauges []string

	length := GetAllFarmingPoolLength(voter)

	for i := 0; i < length; i++ {
		poolAddress, err := voter.Pools(&bind.CallOpts{}, IntToBigInt(i))
		if err != nil {
			panic(err)
		}
		pools = append(pools, poolAddress.Hex())
	}

	for _, poolAddress := range pools {
		gauge, err := voter.Gauges(&bind.CallOpts{}, common.HexToAddress(poolAddress))
		if err != nil {
			panic(err)
		}
		gauges = append(gauges, gauge.Hex())
	}

	return length, gauges
}

func GetAllGauges(voterAddress string) (int, []string) {
	var pools []string
	var gauges []string

	voter := GetVoterInstance(voterAddress)

	length := GetAllFarmingPoolLength(voter)

	for i := 0; i < length; i++ {
		poolAddress, err := voter.Pools(&bind.CallOpts{}, IntToBigInt(i))
		if err != nil {
			panic(err)
		}
		pools = append(pools, poolAddress.Hex())
	}

	for _, poolAddress := range pools {
		gauge, err := voter.Gauges(&bind.CallOpts{}, common.HexToAddress(poolAddress))
		if err != nil {
			panic(err)
		}
		gauges = append(gauges, gauge.Hex())
	}

	return length, gauges
}

