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

# stop exporter to check alerting
docker stop exporter

# check alert status
open http://localhost:9090/alerts

# check alertmanager
open http://localhost:9093/#/alerts

# check 
docker-compose down
```