# Operators
- aggregaration opearator (avg, sum, count, etc.)
- logical operator (AND,OR,UNLESS)
- binary opearator 
  - arithmetic operators (+,-,%,*, etc.)
  - comparison operators (==, <, >, >=, etc.)

# Aggregation Operator

## Grouping sum() with by() and without()
```bash
# call without groupting
PROM_QUERY='prometheus_http_requests_total{handler=~"/status|/alerts"}'

curl 'https://demo.promlabs.com/api/v1/query' \
--data query=$PROM_QUERY | jq

{
  "status": "success",
  "data": {
    "resultType": "vector",
    "result": [
      {
        "metric": {
          "__name__": "prometheus_http_requests_total",
          "code": "200",
          "handler": "/alerts",
          "instance": "prometheus:9090",
          "job": "prometheus"
        },
        "value": [
          1688818950.643,
          "27"
        ]
      },
      {
        "metric": {
          "__name__": "prometheus_http_requests_total",
          "code": "200",
          "handler": "/status",
          "instance": "prometheus:9090",
          "job": "prometheus"
        },
        "value": [
          1688818950.643,
          "14"
        ]
      }
    ]
  }
}

# call with groupting with by()
PROM_QUERY='sum(prometheus_http_requests_total{handler=~"/status|/alerts"})by(handler)'

curl 'https://demo.promlabs.com/api/v1/query' \
--data query=$PROM_QUERY | jq

{
  "status": "success",
  "data": {
    "resultType": "vector",
    "result": [
      {
        "metric": {
          "handler": "/alerts"
        },
        "value": [
          1688819243.856,
          "27"
        ]
      },
      {
        "metric": {
          "handler": "/status"
        },
        "value": [
          1688819243.856,
          "14"
        ]
      }
    ]
  }
}

# call with groupting with without()
PROM_QUERY='sum(prometheus_http_requests_total{handler=~"/status|/alerts"})without(code, instance, job)'

curl 'https://demo.promlabs.com/api/v1/query' \
--data query=$PROM_QUERY | jq

# same result as before from grouping with by() function

{
  "status": "success",
  "data": {
    "resultType": "vector",
    "result": [
      {
        "metric": {
          "handler": "/alerts"
        },
        "value": [
          1688819243.856,
          "27"
        ]
      },
      {
        "metric": {
          "handler": "/status"
        },
        "value": [
          1688819243.856,
          "14"
        ]
      }
    ]
  }
}

```

## Logical Operator

### OR
```bash
topk(2,go_gc_duration_seconds) OR topk(2,prometheus_http_requests_total)
```
 
### AND
```bash
topk(2,go_gc_duration_seconds) AND topk(2,prometheus_http_requests_total)
```
### UNLESS
```bash
topk(2,go_gc_duration_seconds) UNLESS topk(2,prometheus_http_requests_total)
```

## Binary

### Arithmetic Operator
```bash
avg(go_gc_duration_seconds) + 999
```

### Comparison Operator
```bash
avg(go_gc_duration_seconds) < 100000
```