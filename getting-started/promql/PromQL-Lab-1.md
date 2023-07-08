# Vector

- demo_api_http_requests_in_progress = Counter Metric

```bash
PROM_QUERY='demo_api_http_requests_in_progress'

curl 'https://demo.promlabs.com/api/v1/query' \
--data query=$PROM_QUERY | jq

{
  "status": "success",
  "data": {
    "resultType": "vector",
    "result": [
      {
        "metric": {
          "__name__": "demo_api_http_requests_in_progress",
          "instance": "demo-service-0:10000",
          "job": "demo"
        },
        "value": [
          1688797295.464,
          "1"
        ]
      },
      {
        "metric": {
          "__name__": "demo_api_http_requests_in_progress",
          "instance": "demo-service-1:10001",
          "job": "demo"
        },
        "value": [
          1688797295.464,
          "1"
        ]
      },
      {
        "metric": {
          "__name__": "demo_api_http_requests_in_progress",
          "instance": "demo-service-2:10002",
          "job": "demo"
        },
        "value": [
          1688797295.464,
          "0"
        ]
      }
    ]
  }
}
```

# Instat vector
```bash
PROM_QUERY='demo_api_http_requests_in_progress'

curl 'https://demo.promlabs.com/api/v1/query' \
--data query=$PROM_QUERY \
--data time=1688795777| jq

{
  "status": "success",
  "data": {
    "resultType": "vector",
    "result": [
      {
        "metric": {
          "__name__": "demo_api_http_requests_in_progress",
          "instance": "demo-service-0:10000",
          "job": "demo"
        },
        "value": [
          1688795777,
          "0"
        ]
      },
      {
        "metric": {
          "__name__": "demo_api_http_requests_in_progress",
          "instance": "demo-service-1:10001",
          "job": "demo"
        },
        "value": [
          1688795777,
          "1"
        ]
      },
      {
        "metric": {
          "__name__": "demo_api_http_requests_in_progress",
          "instance": "demo-service-2:10002",
          "job": "demo"
        },
        "value": [
          1688795777,
          "1"
        ]
      }
    ]
  }
}
```

# Range vector
```bash
PROM_QUERY='demo_api_http_requests_in_progress[30s]'

curl 'https://demo.promlabs.com/api/v1/query' \
--data query=$PROM_QUERY | jq

{
  "status": "success",
  "data": {
    "resultType": "matrix",
    "result": [
      {
        "metric": {
          "__name__": "demo_api_http_requests_in_progress",
          "instance": "demo-service-0:10000",
          "job": "demo"
        },
        "values": [
          [
            1688795759.44,
            "1"
          ],
          [
            1688795774.436,
            "0"
          ]
        ]
      },
      {
        "metric": {
          "__name__": "demo_api_http_requests_in_progress",
          "instance": "demo-service-1:10001",
          "job": "demo"
        },
        "values": [
          [
            1688795750.233,
            "1"
          ],
          [
            1688795765.233,
            "1"
          ]
        ]
      },
      {
        "metric": {
          "__name__": "demo_api_http_requests_in_progress",
          "instance": "demo-service-2:10002",
          "job": "demo"
        },
        "values": [
          [
            1688795752.852,
            "0"
          ],
          [
            1688795767.852,
            "1"
          ]
        ]
      }
    ]
  }
}
```
