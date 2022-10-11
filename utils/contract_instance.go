package utils

import (
	"dddx-lp-data/abigen/factory"
	"dddx-lp-data/abigen/gauge"
	"dddx-lp-data/abigen/pair"
	"dddx-lp-data/abigen/token"
	"dddx-lp-data/abigen/ve"
	"dddx-lp-data/abigen/voter"
	"dddx-lp-data/initializers"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var DDDXFactory *factory.Factory
var DDDXVeNFT *ve.Ve
var DDDXVoter *voter.Voter

func InitializeContractInstance() {
	DDDXFactory = GetFactoryInstance("0xb5737A06c330c22056C77a4205D16fFD1436c81b")
	DDDXVeNFT = GetVeInstance("0xFe9e21e78089094E1443169c4c74bBBBcBb13DE0")
	DDDXVoter = GetVoterInstance("0xAd8Ab2C2270Ab0603CFC674d28fd545495369f31")
}

// GetTokenInstance generates a new erc20 contract instance.
func GetTokenInstance(address string) *token.Token {
	token, err := token.NewToken(common.HexToAddress(address), initializers.Client)
	if err != nil {
		panic(err)
	}

	return token
}

// GetTokenInstance generates a new erc20 contract instance.
func GetFactoryInstance(address string) *factory.Factory {
	factory, err := factory.NewFactory(common.HexToAddress(address), initializers.Client)
	if err != nil {
		panic(err)
	}

	return factory
}

// GetPoolInstance generates a new pool contract instance.
func GetPoolInstance(address string) *pair.Pair {
	pool, err := pair.NewPair(common.HexToAddress(address), initializers.Client)
	if err != nil {
		panic(err)
	}

	return pool
}

// GetGaugeInstance generates a new gauge contract instance.
func GetGaugeInstance(address string) *gauge.Gauge {
	gauge, err := gauge.NewGauge(common.HexToAddress(address), initializers.Client)
	if err != nil {
		panic(err)
	}

	return gauge
}

// GetVeInstance generates a new vote-escrowed-nft contract instance.
func GetVeInstance(address string) *ve.Ve {
	veNFT, err := ve.NewVe(common.HexToAddress(address), initializers.Client)
	if err != nil {
		panic(err)
	}

	return veNFT
}

func GetVoterInstance(address string) *voter.Voter {
	voter, err := voter.NewVoter(common.HexToAddress(address), initializers.Client)
	if err != nil {
		panic(err)
	}

	return voter
}

// GetTokenPairInstances generates two new erc20 contract instances from pool Instance.
func GetTokenPairInstances(poolInstance *pair.Pair) (*token.Token, *token.Token) {
	token0Address, token1Address, err := poolInstance.Tokens(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	return GetTokenInstance(token0Address.String()), GetTokenInstance(token1Address.String())
}
