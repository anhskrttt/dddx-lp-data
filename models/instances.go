package models

var PairIds []string = []string{"bsc-dddx-wbnb_usd"}
var TokenIds []string = []string{"bnb-token", "wbnb-token", "busd-token", "usdt-token"}
var Tokens []Token = []Token{BNBToken, WBNBToken, BUSDToken, USDTToken}

var User0 = User{
	Address: "0x043220ac21c0ce367689d93822ad70fe95ea8d2e",
	Protocols: []Protocol{
		DDDXProtocol,
	},
}

var User1 User
var User2 User

var BNBToken Token
var WBNBToken Token
var BUSDToken Token
var USDTToken Token

var BSCChain Chain
var DDDXProtocol Protocol

func InitializeModelInstance() {

	BNBToken := Token{
		ID:       "bnb-token",
		Name:     "BNB",
		Symbol:   "BNB",
		Decimals: 18,
	}

	WBNBToken = Token{
		ID:       "wbnb-token",
		Name:     "WBNB",
		Symbol:   "WBNB",
		Decimals: 18,
	}

	BUSDToken = Token{
		ID:       "busd-token",
		Name:     "Binance USD",
		Symbol:   "BUSD",
		Decimals: 18,
	}

	USDTToken = Token{
		ID:       "usdt-token",
		Name:     "Tether",
		Symbol:   "USDT",
		Decimals: 18,
	}

	BSCChain = Chain{
		ID:          "bsc-chain",
		Name:        "BNB Smart Chain",
		NativeToken: BNBToken,
		URL:         "https://www.bnbchain.org/en/smartChain",
		Explorer:    "https://bscscan.com/",
		RPC:         []string{"https://bsc-dataseed1.binance.org/", "https://bsc-dataseed2.binance.org/", "https://bsc-dataseed3.binance.org/"},
		Faucet:      []string{},
		ChainId:     "",
		NetworkId:   "",
	}

	DDDXProtocol = Protocol{
		ID:          "dddx-protocol",
		Name:        "DDDX",
		Symbol:      "DDDX",
		Address:     "0x4B6ee8188d6Df169E1071a7c96929640D61f144f",
		Chain:       BSCChain,
		URL:         "https://dddx.io/",
		Description: "DDDX protocol natively supports swaps between closely correlated assets via a new curve，which realizes extremely low slippage close to 0 and a low exchange fee rate 0.01%. DDDX protocol supports veNFT governance and ve(3,3) tokenomics.",
		Gecko_ID:    "dddx-protocol",
		Tags:        []string{"DEX"},
		TVL:         263645,
	}
}
