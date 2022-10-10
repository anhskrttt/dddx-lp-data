---
description: Get data related to the defi protocol.
---

# Protocol

## Get All Liquidity Pools

{% swagger method="get" path="/api/v1/user/all_pools" baseUrl="https://localhost:3000" summary="Get all liquidity pool of protocol." %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="query" name="protocol_id" type="string" %}
Default is "dddx"
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="Sucessful" %}
```json
{
    "response": {
        "protocol_id": "dddx",
        "length": 57,
        "pools": [
            "0x53A02ad8413293b46153EA447661c7aBC0778292",
            "0x0e04b12B7BbDdEd49e2C0d8D68f323e049AaD08D",
            "0xB4434F5C92Bb7c0647fbC3486b076016A43A2A36",
            "0x88Db6A6b92641E9aC28F4ac80D02A634D14fba59",
            "0x24f1240498BC2d958C57dB65f2785df495997002",
            "0x322Ba943c19f9ec1EF8ceAD8260b30E789Ca1846",
            "0x7dB9c49611c2D749146234eCdA163E83F33Abea1",
            "0x96F9546F188425A126e31cf3da358Bc0eD683EBb",
            // ...
        ]
    }
}
```
{% endswagger-response %}
{% endswagger %}
