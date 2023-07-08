# Summary

- go_gc_duration_seconds = Summary Metric


# Vector
```bash
PROM_QUERY='go_gc_duration_seconds{job="node"}'

curl 'https://demo.promlabs.com/api/v1/query' \
--data query=$PROM_QUERY | jq

{
  "status": "success",
  "data": {
    "resultType": "vector",
    "result": [
      {
        "metric": {
          "__name__": "go_gc_duration_seconds",
          "instance": "node-exporter:9100",
          "job": "node",
          "quantile": "0"
        },
        "value": [
          1688802351.949,
          "0.000046627"
        ]
      },
      {
        "metric": {
          "__name__": "go_gc_duration_seconds",
          "instance": "node-exporter:9100",
          "job": "node",
          "quantile": "0.25"
        },
        "value": [
          1688802351.949,
          "0.000053263"
        ]
      },
      {
        "metric": {
          "__name__": "go_gc_duration_seconds",
          "instance": "node-exporter:9100",
          "job": "node",
          "quantile": "0.5"
        },
        "value": [
          1688802351.949,
          "0.000070411"
        ]
      },
      {
        "metric": {
          "__name__": "go_gc_duration_seconds",
          "instance": "node-exporter:9100",
          "job": "node",
          "quantile": "0.75"
        },
        "value": [
          1688802351.949,
          "0.000083544"
        ]
      },
      {
        "metric": {
          "__name__": "go_gc_duration_seconds",
          "instance": "node-exporter:9100",
          "job": "node",
          "quantile": "1"
        },
        "value": [
          1688802351.949,
          "0.001591335"
        ]
      }
    ]
  }
}
```

# Quantile 0.5

- 50% of go_gc_duration_seconds is faster than '0.000070411'

```bash
PROM_QUERY='go_gc_duration_seconds{job="node",quantile="0.5"}'

curl 'https://demo.promlabs.com/api/v1/query' \
--data query=$PROM_QUERY | jq

{
  "status": "success",
  "data": {
    "resultType": "vector",
    "result": [
      {
        "metric": {
          "__name__": "go_gc_duration_seconds",
          "instance": "node-exporter:9100",
          "job": "node",
          "quantile": "0.5"
        },
        "value": [
          1688802463.927,
          "0.000070411"
        ]
      }
    ]
  }
}
``````
