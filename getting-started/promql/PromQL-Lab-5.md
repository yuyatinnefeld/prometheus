# Histogram

- go_gc_pauses_seconds_total = Histogram Metric


# Histogram Count
```bash
PROM_QUERY='go_gc_pauses_seconds_total_bucket{instance="demo-service-0:10000"}'

curl 'https://demo.promlabs.com/api/v1/query' \
--data query=$PROM_QUERY | jq

{
  "status": "success",
  "data": {
    "resultType": "vector",
    "result": [
      {
        "metric": {
          "__name__": "go_gc_pauses_seconds_total_bucket",
          "instance": "demo-service-0:10000",
          "job": "demo",
          "le": "+Inf"
        },
        "value": [
          1688801821.131,
          "9793034"
        ]
      },
      {
        "metric": {
          "__name__": "go_gc_pauses_seconds_total_bucket",
          "instance": "demo-service-0:10000",
          "job": "demo",
          "le": "-5e-324"
        },
        "value": [
          1688801821.131,
          "0"
        ]
      },
      {
        "metric": {
          "__name__": "go_gc_pauses_seconds_total_bucket",
          "instance": "demo-service-0:10000",
          "job": "demo",
          "le": "0.00016383999999999998"
        },
        "value": [
          1688801821.131,
          "9200370"
        ]
      },
      {
        "metric": {
          "__name__": "go_gc_pauses_seconds_total_bucket",
          "instance": "demo-service-0:10000",
          "job": "demo",
          "le": "0.0020971519999999997"
        },
        "value": [
          1688801821.131,
          "9682876"
        ]
      },
      {
        "metric": {
          "__name__": "go_gc_pauses_seconds_total_bucket",
          "instance": "demo-service-0:10000",
          "job": "demo",
          "le": "0.020971519999999997"
        },
        "value": [
          1688801821.131,
          "9792373"
        ]
      },
      {
        "metric": {
          "__name__": "go_gc_pauses_seconds_total_bucket",
          "instance": "demo-service-0:10000",
          "job": "demo",
          "le": "0.26843545599999996"
        },
        "value": [
          1688801821.131,
          "9793018"
        ]
      },
      {
        "metric": {
          "__name__": "go_gc_pauses_seconds_total_bucket",
          "instance": "demo-service-0:10000",
          "job": "demo",
          "le": "1.2799999999999998e-06"
        },
        "value": [
          1688801821.131,
          "35"
        ]
      },
      {
        "metric": {
          "__name__": "go_gc_pauses_seconds_total_bucket",
          "instance": "demo-service-0:10000",
          "job": "demo",
          "le": "1.2799999999999998e-07"
        },
        "value": [
          1688801821.131,
          "0"
        ]
      },
      {
        "metric": {
          "__name__": "go_gc_pauses_seconds_total_bucket",
          "instance": "demo-service-0:10000",
          "job": "demo",
          "le": "1.6383999999999998e-05"
        },
        "value": [
          1688801821.131,
          "2763870"
        ]
      },
      {
        "metric": {
          "__name__": "go_gc_pauses_seconds_total_bucket",
          "instance": "demo-service-0:10000",
          "job": "demo",
          "le": "9.999999999999999e-09"
        },
        "value": [
          1688801821.131,
          "0"
        ]
      },
      {
        "metric": {
          "__name__": "go_gc_pauses_seconds_total_bucket",
          "instance": "demo-service-0:10000",
          "job": "demo",
          "le": "9.999999999999999e-10"
        },
        "value": [
          1688801821.131,
          "0"
        ]
      }
    ]
  }
}
```

# Histogram Count
```bash
PROM_QUERY='go_gc_pauses_seconds_total_count{instance="demo-service-0:10000"}'

curl 'https://demo.promlabs.com/api/v1/query' \
--data query=$PROM_QUERY | jq

{
  "status": "success",
  "data": {
    "resultType": "vector",
    "result": [
      {
        "metric": {
          "__name__": "go_gc_pauses_seconds_total_count",
          "instance": "demo-service-0:10000",
          "job": "demo"
        },
        "value": [
          1688801924.925,
          "9793104"
        ]
      }
    ]
  }
}
```

## 0.9 Quantile

- 90% of go_gc_pauses_second is faster than '0.00014239185454545448'

```bash
PROM_QUERY='histogram_quantile(0.9,rate(go_gc_pauses_seconds_total_bucket{instance="demo-service-0:10000"}[1m]))'

curl 'https://demo.promlabs.com/api/v1/query' \
--data query=$PROM_QUERY | jq

{
  "status": "success",
  "data": {
    "resultType": "vector",
    "result": [
      {
        "metric": {
          "instance": "demo-service-0:10000",
          "job": "demo"
        },
        "value": [
          1688802078.515,
          "0.00014239185454545448"
        ]
      }
    ]
  }
}
```