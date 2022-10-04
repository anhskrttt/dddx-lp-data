# Quick Start



{% hint style="info" %}
**Good to know:** This is an in-development project. Don't forget to checkout the incoming APIs!
{% endhint %}

## Clone project

The best way to interact with our API is to use one of our official libraries:

{% tabs %}
{% tab title="Bash" %}
```git
git clone https://github.com/anhskrttt/dddx-lp-data.git
cd $_
```
{% endtab %}
{% endtabs %}

## Install dependencies

To make your first request, send an authenticated request to the pets endpoint. This will create a `pet`, which is nice.

{% tabs %}
{% tab title="Bash" %}
```bash
go mod tidy
```
{% endtab %}
{% endtabs %}

## Run project

{% tabs %}
{% tab title="Bash" %}
```bash
CompileDaemon -command="./dddx-lp-data"
```
{% endtab %}
{% endtabs %}

## Access GoDoc

{% tabs %}
{% tab title="Bash" %}
```bash
godoc -http=:6060
```
{% endtab %}
{% endtabs %}
