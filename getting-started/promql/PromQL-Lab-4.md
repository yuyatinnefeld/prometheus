# rate()

- promhttp_metric_handler_requests_total = Counter Metric

# Vector (total requests of status 200 from prom handler)

```bash
PROM_QUERY='promhttp_metric_handler_requests_total{instance="node-exporter:9100",code="200"}'

curl 'https://demo.promlabs.com/api/v1/query' \
--data query=$PROM_QUERY | jq

{
  "status": "success",
  "data": {
    "resultType": "vector",
    "result": [
      {
        "metric": {
          "__name__": "promhttp_metric_handler_requests_total",
          "code": "200",
          "instance": "node-exporter:9100",
          "job": "node"
        },
        "value": [
          1688799375.177,
          "952594"
        ]
      }
    ]
  }
}
```

# Range Vector

```bash
PROM_QUERY='promhttp_metric_handler_requests_total{instance="node-exporter:9100",code="200"}[1m]'

curl 'https://demo.promlabs.com/api/v1/query' \
--data query=$PROM_QUERY | jq

{
  "status": "success",
  "data": {
    "resultType": "matrix",
    "result": [
      {
        "metric": {
          "__name__": "promhttp_metric_handler_requests_total",
          "code": "200",
          "instance": "node-exporter:9100",
          "job": "node"
        },
        "values": [
          [
            1688799602.944,
            "952610"
          ],
          [
            1688799617.944,
            "952611"
          ],
          [
            1688799632.944,
            "952612"
          ],
          [
            1688799647.944,
            "952613"
          ]
        ]
      }
    ]
  }
}
```

# increase()

- requests are increased '4'

```bash
PROM_QUERY='increase(promhttp_metric_handler_requests_total{instance="node-exporter:9100",code="200"}[1m])'

curl 'https://demo.promlabs.com/api/v1/query' \
--data query=$PROM_QUERY | jq

{
  "status": "success",
  "data": {
    "resultType": "vector",
    "result": [
      {
        "metric": {
          "code": "200",
          "instance": "node-exporter:9100",
          "job": "node"
        },
        "value": [
          1688799972.254,
          "4"
        ]
      }
    ]
  }
}
```

# Rate()
- per-second average rate of change within 1 min
- requests increased => 4
- rate calculates per-second => 60s
- 4 req / 60 sec = 0.06666666666666665

```bash
PROM_QUERY='rate(promhttp_metric_handler_requests_total{instance="node-exporter:9100",code="200"}[1m])'

curl 'https://demo.promlabs.com/api/v1/query' \
--data query=$PROM_QUERY | jq

{
  "status": "success",
  "data": {
    "resultType": "vector",
    "result": [
      {
        "metric": {
          "code": "200",
          "instance": "node-exporter:9100",
          "job": "node"
        },
        "value": [
          1688799509.531,
          "0.06666666666666665"
        ]
      }
    ]
  }
}
```