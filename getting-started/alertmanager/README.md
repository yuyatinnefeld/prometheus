# Alertmanager

## Docs
- https://sysdig.com/blog/prometheus-alertmanager/
- https://developers.soundcloud.com/blog/alerting-on-slos
- https://docs.search-guard.com/latest/elasticsearch-alerting-throttling

## Setup
```bash
# add alertmanager prometheus.yml
vi prometheus.yml

# add alertmanager server
vi docker-compose.yml

# start all servers
docker-compose up -d

# check node app
open http://localhost:9100

# check golang app
open http://localhost:8888

# check python app
open http://localhost:8080
curl http://localhost:8080/summary

# check prometheus app
open http://localhost:9090

# stop exporter and python app
docker stop exporter
docker stop my-app-python

# check alert status (prometheus)
open http://localhost:9090/alerts

# check alertmanager
open http://localhost:9093/#/alerts

# check 
docker-compose down
```

