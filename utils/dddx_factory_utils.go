package utils

import (
	"dddx-lp-data/abigen/factory"
	"dddx-lp-data/abigen/gauge"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// GetAllLiquidityPools returns a string array containing addresses of all supported pools of protocol.
func GetAllLiquidityPoolsOf(factoryInstance *factory.Factory) (int, []string) {
	var pools []string

	poolLength := GetAllPairsLength(factoryInstance)

	for i := 0; i < poolLength; i++ {
		poolAddress, err := factoryInstance.AllPairs(&bind.CallOpts{}, IntToBigInt(i))
		if err != nil {
			panic(err)
		}
		pools = append(pools, poolAddress.Hex())
	}

	return poolLength, pools
}

// GetAllLiquidityPools returns a string array containing addresses of all supported pools of protocol.
func GetAllLiquidityPools(factoryAddress string) (int, []string) {
	var pools []string

	poolLength := GetAllPairsLength(DDDXFactory)

	for i := 0; i < poolLength; i++ {
		poolAddress, err := DDDXFactory.AllPairs(&bind.CallOpts{}, IntToBigInt(i))
		if err != nil {
			panic(err)
		}
		pools = append(pools, poolAddress.Hex())
	}

	return poolLength, pools
}

func GetAllPairsLength(factoryInstance *factory.Factory) int {
	poolLength, err := factoryInstance.AllPairsLength(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	return int(poolLength.Uint64())
}

func GetFarmBalFromGauge(userAddress string, gaugeInstance *gauge.Gauge) *big.Int {
	lpBal, err := gaugeInstance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(userAddress))
	if err != nil {
		panic(err)
	}

	return lpBal
}

func GetFarmBalFromGaugeAddress(userAddress string, gaugeAddress string) *big.Int {
	gaugeInstance := GetGaugeInstance(gaugeAddress)

	lpBal, err := gaugeInstance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(userAddress))
	if err != nil {
		panic(err)
	}

	return lpBal
}
