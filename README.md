# DDDX Liquidity Provider Data
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
| /api/:usrAddress      | GET           | None              | Get wallet address's balance information  |

## Resources
* [DDDX Contracts](https://dddx.gitbook.io/dddx.io/tokenomics/contracts)