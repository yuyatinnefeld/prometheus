# avg() vs avg_over_time()

- go_gc_duration_seconds = Summary Metric

# Calc AVG with avg()

- aggregates avg value of all quantiles

```bash
PROM_QUERY='avg(go_gc_duration_seconds{job="node"})'

curl 'https://demo.promlabs.com/api/v1/query' \
--data query=$PROM_QUERY | jq

{
  "status": "success",
  "data": {
    "resultType": "vector",
    "result": [
      {
        "metric": {},
        "value": [
          1688798254.276,
          "0.0002463284"
        ]
      }
    ]
  }
}
```

# Calc AVG with avg_over_time()

- calculates avg value pre quantile

```bash
PROM_QUERY='avg_over_time(go_gc_duration_seconds{job="node"}[1m])'

curl 'https://demo.promlabs.com/api/v1/query' \
--data query=$PROM_QUERY | jq

{
  "status": "success",
  "data": {
    "resultType": "vector",
    "result": [
      {
        "metric": {
          "instance": "node-exporter:9100",
          "job": "node",
          "quantile": "0"
        },
        "value": [
          1688798447.349,
          "0.000046145"
        ]
      },
      {
        "metric": {
          "instance": "node-exporter:9100",
          "job": "node",
          "quantile": "0.25"
        },
        "value": [
          1688798447.349,
          "0.0000523145"
        ]
      },
      {
        "metric": {
          "instance": "node-exporter:9100",
          "job": "node",
          "quantile": "0.5"
        },
        "value": [
          1688798447.349,
          "0.000059885"
        ]
      },
      {
        "metric": {
          "instance": "node-exporter:9100",
          "job": "node",
          "quantile": "0.75"
        },
        "value": [
          1688798447.349,
          "0.00007966350000000001"
        ]
      },
      {
        "metric": {
          "instance": "node-exporter:9100",
          "job": "node",
          "quantile": "1"
        },
        "value": [
          1688798447.349,
          "0.000992408"
        ]
      }
    ]
  }
}
```