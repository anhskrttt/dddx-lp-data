# DDDX Liquidity Provider Data
## Document Reference
[DDDX-LP-DATA](https://anhs-organization.gitbook.io/dddx-data/)
## Structure
```
$ tree
.
├── abigen                              // All go files gen by abigen
│   ├── factory
│   │   └── factory.go
│   ├── ...
├── build                               // abi files gen by solc
│   ├── Base64.abi
│   ├── BaseV1Factory.abi
│   ├── ...
├── contracts                           // All contracts used
│   ├── BaseV1Factory.sol
│   ├── ...
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
├── routers                             // Routers used
│   └── routers.go
└── utils                               // Util functions
    ├── contract_instance.go
    ├── contract_utils.go
    ├── dddx_details.go
    ├── dddx_erc20_utils.go
    ├── dddx_factory_utils.go
    ├── dddx_gauge_utils.go
    ├── dddx_pair_utils.go
    ├── dddx_voter_utils.go
    ├── local_storage.go
    ├── math.go
    └── token.go

15 directories, 58 files
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