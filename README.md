# whoamip

Whoamip provides a HTTP endpoint that returns your current IP address.

## Usage

```$
$ whoami
Usage of whoamip:
  -metricsPort int
        Metrics Listening Port (default 9100)
  -port int
        Listening Port (default 8080)

$ whoamip
INFO[0000] Listening on tcp://0.0.0.0:8080
INFO[0000] Metrics listening on tcp://0.0.0.0:9100
```

## Example

```$
$ curl http://192.168.1.51:8080 | jq
{
  "ip": "192.168.1.109"
}
```
