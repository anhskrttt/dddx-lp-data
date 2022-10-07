# DDDX Liquidity Provider Data
## Document Reference
URL: [DDDX-LP-DATA](https://anhs-organization.gitbook.io/dddx-data/)
## Structure
```
.
├── abigen                              // All go files gen by abigen
│   ├── gauge
│   │   └── gauge.go
│   ├── pair
│   │   └── pair.go
│   ├── token
│   │   └── token.go
│   └── ve
│       └── ve.go
├── build                               // abi files gen by solc
│   ├── Base64.abi
│   ├── BaseV1Factory.abi
│   ├── BaseV1Fees.abi
│   ├── BaseV1GaugeFactory.abi
│   ├── BaseV1Pair.abi
│   ├── erc20.abi
│   ├── Gauge.abi
│   ├── IBaseV1Callee.abi
│   ├── IBaseV1Core.abi
│   ├── IBaseV1Factory.abi
│   ├── IBribe.abi
│   ├── IERC165.abi
│   ├── IERC20.abi
│   ├── IERC721.abi
│   ├── IERC721Metadata.abi
│   ├── IERC721Receiver.abi
│   ├── Math.abi
│   └── ve.abi
├── contracts                           // All contracts used
│   ├── BaseV1Factory.sol
│   ├── BaseV1GaugeFactory.sol
│   ├── BaseV1Router01.sol
│   └── ve.sol
├── controllers                         // Handler functions
│   └── controllers.go
├── dddx-lp-data                        // CompileDaemon running file
├── go.mod
├── go.sum
├── initializers                        // Initializations
│   ├── connect_to_eth_client.go
│   └── load_env.go
├── main.go                             // Main file
├── models                              // All models used
│   ├── models.go
│   ├── request.go
│   └── response.go
├── README.md
├── routers
│   └── routers.go                      // Routers used
└── utils                               // Util functions
    ├── contract_utils.go
    ├── math.go
    └── utils.go

12 directories, 41 files
```
## Setups
- Clone project
```
git clone https://github.com/anhskrttt/dddx-lp-data
```
- Create .env file with proper content
```
cp .env.example .env
```
- Run project
```
go mod tidy
CompileDaemon -command="./dddx-lp-data"
```


## Routes
| Routes                | HTTP Methods  | Parameters        | Description                               |
| -------------         | ------------- | ----------------- | ----------------------------------------- |
| api/v1/user/lps      | GET           | user_address, protocol_id, gauge_address              | Get wallet address's LP balance  |
| api/v1/user/farms      | GET           | user_address, protocol_id, gauge_address              | Get wallet address's Farming LP balance  |
| api/v1/user/staked      | GET           | user_address, protocol_id vote_escrowed_address            | Get wallet address's staked balance   |
| api/v1/user/all_pools      | GET           | user_address, protocol_id           | Get wallet address's all LP balance   |
| api/v1/user/all_farms      | GET           | user_address, protocol_id           | Get wallet address's all farming balance   |
| api/v1/user      | GET           | user_address, protocol_id           | Get wallet address's all LP, farming and staked balance   |

## Resources
* [DDDX Contracts](https://dddx.gitbook.io/dddx.io/tokenomics/contracts)