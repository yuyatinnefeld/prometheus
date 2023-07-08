# Operators
- aggregaration opearator
  ```bash
  sum (calculate sum over dimensions)
  min (select minimum over dimensions)
  max (select maximum over dimensions)
  avg (calculate the average over dimensions)
  group (all values in the resulting vector are 1)
  stddev (calculate population standard deviation over dimensions)
  stdvar (calculate population standard variance over dimensions)
  count (count number of elements in the vector)
  count_values (count number of elements with the same value)
  bottomk (smallest k elements by sample value)
  topk (largest k elements by sample value)
  quantile (calculate φ-quantile (0 ≤ φ ≤ 1) over dimensions)
  ```

  # examples
  COUNT(prometheus_http_requests_total{code=~"2.*|4.*"}) BY (code)
  MAX(prometheus_http_requests_total{code=~"4.*"}) BY (handler)
  AVG(prometheus_http_requests_total) BY (code)
  # 
  ```
- logical/set operator
  ```bash
  # intersection
  AND
  # union
  OR
  # complement
  UNLESS
  ```
- binary opearator 
  - arithmetic operators
    ```bash
    + (add)
    – (subtract)
    * (multiply)
    / (divide)
    % (percentage)
    ^ (exponents)
    ```
  - comparison operators
    ```bash
    == (equal to)
    != (does not equal)
    > (greater than)
    < (less than)
    >= (greater than or equal to)
    <= (less than or equal to)
    ```
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