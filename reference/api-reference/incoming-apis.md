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

{% swagger-parameter in="query" name="user_address" type="string" %}

{% endswagger-parameter %}

{% swagger-parameter in="query" name="protocol_id" type="string" %}
All protocol ids supported can be found in GET /api/v1/all_protocol
{% endswagger-parameter %}
{% endswagger %}

{% swagger method="get" path="/api/v1/all_protocol" baseUrl="https://localhost:3000" summary="Get All Protocol Supported" %}
{% swagger-description %}
Updating...
{% endswagger-description %}
{% endswagger %}
