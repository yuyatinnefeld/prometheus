# Average

- demo_api_http_requests_in_progress = Counter Metric

# Calc Average with avg()
```bash
PROM_QUERY='avg(demo_api_http_requests_in_progress)'

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
          1688797481.314,
          "1.6666666666666667"
        ]
      }
    ]
  }
}
```

# Calc Average with sum() / count()
```bash
PROM_QUERY='sum(demo_api_http_requests_in_progress)/count(demo_api_http_requests_in_progress)'

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
          1688797481.314,
          "1.6666666666666667"
        ]
      }
    ]
  }
}
```