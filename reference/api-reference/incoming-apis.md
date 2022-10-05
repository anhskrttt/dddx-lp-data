---
description: >-
  APIs document for dddx-lp-data project:
  https://github.com/anhskrttt/dddx-lp-data
---

# Incoming APIs

## Get User LP balance

{% swagger method="get" path="/api/v1/user/lps" baseUrl="https://localhost:3000" summary="Get All Pools in Protocol" %}
{% swagger-description %}
Updating...
{% endswagger-description %}

{% swagger-parameter in="query" name="protocol_id" type="string" %}
All protocol ids supported can be found in GET /api/v1/all_protocol
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="Successful" %}
```json
{
    "pools": [
        "0x53A02ad8413293b46153EA447661c7aBC0778292",
        "0x0e04b12B7BbDdEd49e2C0d8D68f323e049AaD08D",
        "0xB4434F5C92Bb7c0647fbC3486b076016A43A2A36",
        "0x88Db6A6b92641E9aC28F4ac80D02A634D14fba59",
        "0x24f1240498BC2d958C57dB65f2785df495997002",
        "0x322Ba943c19f9ec1EF8ceAD8260b30E789Ca1846",
        "0x7dB9c49611c2D749146234eCdA163E83F33Abea1",
        "0x96F9546F188425A126e31cf3da358Bc0eD683EBb",
        "0xB76AB2367Cf13A4D63A2c9bE760c6703Bf1B87EC",
        "0x30e51d84Fd4CFDd8a4D1B3Cc7BcA3DFeFeb39937",
        "0x08C13019b0AFcF0c43bE868761bb1DF69dB8CDF1",
        "0x27d1ADf205872b13E9091d3E5EFCe6D975e3c1c4",
        "0x9Fa51CC5a51Bc61f7668f753C1A85DF0fa87C7FF",
        "0x12706deA61D9789c77AC8CBC4b2De98f86c31210",
        "0xe6867A7fEa57b585Dd221d8249d07709829F2868",
        "0xcea9b56953F60172042dC37275A82fC948350B1B",
        "0xc788EDBBC45848243914EB9F910502031A39DD95",
        "0xc005f05E45Bd3035F49405B0BbC109bD8504980b",
        "0x4Eb44B822A80E1B6F8cF242543b9EC6a7a4D42D8",
        "0xC530666ACEB3eC9e970fAD8bC2955AcB789a219C",
        "0x3f69e12C01C47F5e00F28653be304662c9B9E976",
        "0xEb9F2feb6e37cA218fDa843712FBAE7C9FAAc97D",
        "0xA151dEf9de9D254fd82Af03191317cA9dF2c7B41",
        "0x33B8A376eeD3E00d97D69bb845D6D7f816EcC4dA",
        "0xa4144746Db07AB1BcB41Dd2164fE20B1e35dB052",
        "0x3e43Bcc889CC4505a68a89990E73cd47B54fCC88",
        "0xc5E0914dad423550c021b6a2706aD1dB954B285F",
        "0x08879C5855Ef7463be4EBD16460D77D4AE576c70",
        "0xf697e149864BECd29ce4edfA0ef93C4fc3121C63",
        "0xf902df5F8d9d7b6BBaB0Dc6eDFBDdc8dD7b3CD41",
        "0xeFb2A8AE74B507683F44774778db2aD265c49f5B",
        "0x99e0aF7aF44b489DD96a9A0CBeD326726613A815",
        "0x310df61165E5D9389B77F0DbfFF1Ec26a6D9F5F5",
        "0x3d3D07572dC2b53276C7ddbeF5719660D94d37Eb",
        "0x20293737d61BB937e9a7d0113a9D23f7A7F6e180",
        "0x505DbB17474ce1d97590799ec6C74644cDF110d3",
        "0x942D81C5E4d4876b7aa653eeAf60245e37F5964D",
        "0xCdb0C81eCC3d82eb1bDbcB13f7C660b11b69Fd32",
        "0x0a99c85877f657d94ABf7B88ff02f98CD3b51B3a",
        "0x4610F97788e856F67B8FCACE70be00F3DFd9fff4",
        "0x96B3619A5d979594A82134B36e84d189C85eb28D",
        "0xF546d6aD0B1E5b3689b393FE2E7C2b38d56E51BE",
        "0xCC86E684d91F5d2e8fa5bE811137394419699A2D",
        "0x9239Db6899d13b874Cf8B1F642bf00d36fAF2742",
        "0xf77a7865B6198d18a75166363926424BE186ba33",
        "0x2Fec59E3d354b177c6Fb744D752BAafbf00D3Df7",
        "0xFE0679A350f0c36b61370235b43C77a77642a828",
        "0x80643fb78877C3147A71fad7a123465b45f57c81",
        "0x27d49D35B536e65d92ad176d39700Be410fFF781",
        "0x15158711C55D06a0404800fE6520768460dF145b",
        "0x0EBA5590bC0797868c8D9711C7A02aD9527Fd4Ff",
        "0xFA673A51Db01f18A2e6684f956aA23659dF3caa7",
        "0xe04Bc0aC0452c6878914c845eC974cC61b93c77B",
        "0xc0D9C342fe997Cf2a4D7367203019c3fC1Bd6A31",
        "0x0037104b6bB001e2129cAc51874A8D4A6daD61F4",
        "0x9Ec14754Ed86353ab94EF9e7752ca60E3430F3cF",
        "0x7766Add392aC19fAccd8883e4200727ba13F38E4"
    ],
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger method="get" path="/api/v1/all_protocol" baseUrl="https://localhost:3000" summary="Get All Protocol Supported" %}
{% swagger-description %}
Updating...
{% endswagger-description %}
{% endswagger %}
