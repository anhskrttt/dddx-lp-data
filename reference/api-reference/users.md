---
description: >-
  APIs document for dddx-lp-data project:
  https://github.com/anhskrttt/dddx-lp-data
---

# Users

## Get User's All LP balance

{% swagger method="get" path="/api/v1/user/all_pools" baseUrl="https://localhost:3000" summary="Get user's lp balance" %}
{% swagger-description %}
Return all LP balance of user.
{% endswagger-description %}

{% swagger-parameter in="query" name="user_address" type="string" %}

{% endswagger-parameter %}

{% swagger-parameter in="query" name="protocol_id" type="string" required="false" %}
Default is "dddx"
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="Sucessful" %}
```json
{
    "protocol_id": "dddx",
    "response": [
        {
            "token0": {
                "token_symbol": "WBNB",
                "balance": 9762638811166809,
                "balance_in_usd": 2.770832147385364
            },
            "token1": {
                "token_symbol": "BUSD",
                "balance": 2769738641946332084,
                "balance_in_usd": 2.768492259557456
            },
            "pool": {
                "pool_address": "0x322Ba943c19f9ec1EF8ceAD8260b30E789Ca1846",
                "chain_name": "BNB Smart Chain",
                "lp_token": {
                    "name": "VolatileV1 AMM - WBNB/BUSD",
                    "symbol": "vAMM-WBNB/BUSD",
                    "decimals": 18
                }
            }
        }
    ],
    "user_address": "0x043220ac21c0ce367689d93822ad70fe95ea8d2e"
}
```
{% endswagger-response %}
{% endswagger %}

##

## Get User's All Farming balance

{% swagger method="get" path="/api/v1/user/all_farms" baseUrl="https://localhost:3000" summary="Get user's farming balance" %}
{% swagger-description %}
Return all farming balance of user
{% endswagger-description %}

{% swagger-parameter in="query" name="user_address" type="string" %}

{% endswagger-parameter %}

{% swagger-parameter in="query" name="protocol_id" type="string" %}
Default is "dddx"
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="Sucessful" %}
```json
{
    "protocol_id": "dddx",
    "response": [
        {
            "token0": {
                "token_symbol": "DDDX",
                "balance": 35844396179960894696232,
                "balance_in_usd": 91.59748688619567
            },
            "token1": {
                "token_symbol": "WBNB",
                "balance": 325220357759858831,
                "balance_in_usd": 91.68287105608182
            },
            "pool": {
                "pool_address": "0x0e04b12B7BbDdEd49e2C0d8D68f323e049AaD08D",
                "chain_name": "BNB Smart Chain",
                "lp_token": {
                    "name": "VolatileV1 AMM - DDDX/WBNB",
                    "symbol": "vAMM-DDDX/WBNB",
                    "decimals": 18
                }
            }
        },
        {
            "token0": {
                "token_symbol": "DOU",
                "balance": 53803206385688892972486,
                "balance_in_usd": 0
            },
            "token1": {
                "token_symbol": "BUSD",
                "balance": 14987403963791853856,
                "balance_in_usd": 0
            },
            "pool": {
                "pool_address": "0xf77a7865B6198d18a75166363926424BE186ba33",
                "chain_name": "BNB Smart Chain",
                "lp_token": {
                    "name": "VolatileV1 AMM - DOU/BUSD",
                    "symbol": "vAMM-DOU/BUSD",
                    "decimals": 18
                }
            }
        }
    ],
    "user_address": "0x94fac6b9634f00801b122e2c3dfe1c29b44cda25"
}
```
{% endswagger-response %}
{% endswagger %}

## Get user LP balance

{% swagger method="get" path="/api/v1/user/lps" baseUrl="https://localhost:3000" summary="Get LP Data" %}
{% swagger-description %}
LP Data includes balance (in wei) of 02 tokens user provided in pool.
{% endswagger-description %}

{% swagger-parameter in="query" name="user_address" type="string" required="true" %}
User's address
{% endswagger-parameter %}

{% swagger-parameter in="query" name="protocol_id" type="string" required="true" %}
All protocol ids supported can be found in GET /api/v1/all_protocol
{% endswagger-parameter %}

{% swagger-parameter in="query" name="gauge_address" type="string" required="true" %}
Gauge address of LP pool
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="Sucessful" %}
```json
{
    "lp_response": {
        "user_address": "0x043220ac21c0ce367689d93822ad70fe95ea8d2e",
        "token0": {
            "token_symbol": "WBNB",
            "balance": 9699082058579894,
            "balance_in_usd": 2.7760712668067375
        },
        "token1": {
            "token_symbol": "BUSD",
            "balance": 2787859043890862334,
            "balance_in_usd": 2.787859043890862
        },
        "protocol_id": "dddx",
        "pool": {
            "pool_address": "0x322Ba943c19f9ec1EF8ceAD8260b30E789Ca1846",
            "chain_name": "BNB Smart Chain",
            "lp_token": {
                "name": "VolatileV1 AMM - WBNB/BUSD",
                "symbol": "vAMM-WBNB/BUSD",
                "decimals": 18
            }
        }
    }
}
```
{% endswagger-response %}
{% endswagger %}

## Get user farming balance

{% swagger method="get" path="/api/v1/user/farms" baseUrl="https://localhost:3000" summary="Get Farming Data" %}
{% swagger-description %}
LP Data includes balance (in wei) of 02 tokens user's farming.
{% endswagger-description %}

{% swagger-parameter in="query" name="user_address" type="string" %}
User's address
{% endswagger-parameter %}

{% swagger-parameter in="query" name="protocol_id" type="string" %}
All protocol ids supported can be found in GET /api/v1/all_protocol
{% endswagger-parameter %}

{% swagger-parameter in="query" name="gauge_address" type="string" %}
Gauge address of LP pool
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="Sucessful" %}
```json
{
    "farm_response": {
        "user_address": "0x972b0f9cde1266e860e546ac92e783741769400f",
        "token0": {
            "token_symbol": "WBNB",
            "balance": 66683602434827356951,
            "balance_in_usd": 19076.84498455541
        },
        "token1": {
            "token_symbol": "BUSD",
            "balance": 19167224589331462391556,
            "balance_in_usd": 19167.22458933146
        },
        "protocol_id": "dddx",
        "pool": {
            "pool_address": "0x322Ba943c19f9ec1EF8ceAD8260b30E789Ca1846",
            "chain_name": "BNB Smart Chain",
            "lp_token": {
                "name": "VolatileV1 AMM - WBNB/BUSD",
                "symbol": "vAMM-WBNB/BUSD",
                "decimals": 18
            }
        }
    }
}
```
{% endswagger-response %}
{% endswagger %}

## Get user staked balances

{% swagger method="get" path="/api/v1/user/staked" baseUrl="https://localhost:3000" summary="Get Staked Data" %}
{% swagger-description %}
Staked Data includes balance (in wei) of token staked and time to unlock.
{% endswagger-description %}

{% swagger-response status="200: OK" description="Sucessful" %}
```json
{
    "staked_response": {
        "user_address": "0x94fac6b9634f00801b122e2c3dfe1c29b44cda25",
        "token": "DDDX",
        "staked_balance": [
            {
                "Amount": 580293524395114994705,
                "End": 1684368000
            },
            {
                "Amount": 31852303369588166123234,
                "End": 1723075200
            }
        ]
    }
}
```
{% endswagger-response %}
{% endswagger %}
