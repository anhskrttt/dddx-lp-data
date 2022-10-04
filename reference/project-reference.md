---
description: >-
  This page explains core functions in the project. URL:
  https://github.com/anhskrttt/dddx-lp-data
---

# Project Reference

## Structure

```bash
.
├── abigen // 
│   ├── gauge
│   │   └── gauge.go
│   ├── pair
│   │   └── pair.go
│   ├── token
│   │   └── token.go
│   └── ve
│       └── ve.go
├── build
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
├── contracts
│   ├── BaseV1Factory.sol
│   ├── BaseV1GaugeFactory.sol
│   ├── BaseV1Router01.sol
│   └── ve.sol
├── controllers
│   └── controllers.go
├── dddx-lp-data
├── go.mod
├── go.sum
├── initializers
│   ├── connect_to_eth_client.go
│   └── load_env.go
├── main.go
├── models
│   ├── models.go
│   ├── request.go
│   └── response.go
├── README.md
├── routers
│   └── routers.go
└── utils
    ├── contract_utils.go
    ├── math.go
    └── utils.go

12 directories, 41 files
```
