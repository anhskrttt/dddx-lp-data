package utils

import (
	"dddx-lp-data/abigen/gauge"
	"dddx-lp-data/abigen/pair"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func GetPoolAddressFromGaugeAddress(gaugeAddress string) string {
	gaugeInstance := GetGaugeInstance(gaugeAddress)

	// Get pair contract address from gauge contract
	poolAddress, err := gaugeInstance.Stake(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	return poolAddress.Hex()
}

func GetPoolAddressFromGauge(gaugeInstance *gauge.Gauge) string {
	// Get pair contract address from gauge contract
	poolAddress, err := gaugeInstance.Stake(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	return poolAddress.Hex()
}

func GetPoolInstanceFromGauge(gaugeInstance *gauge.Gauge) *pair.Pair {
	// Get pair contract address from gauge contract
	poolAddress, err := gaugeInstance.Stake(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	return GetPoolInstance(poolAddress.Hex())
}
