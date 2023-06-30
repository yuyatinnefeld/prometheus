
# Alerting Rules
```bash
# go alert rules dir
cd alert-rules

# create rules
vi rules.yml

# update prometheus file
vi prometheus.yml

# update docker-compose file
vi docker-compose.yml

# restart docker
docker-compose up -d

# verify alert
open http://localhost:9090/alerts

# test alert
docker stop exporter
open http://localhost:9090/alerts
# Status: INACTIVE > FIRING

# stop
docker-compose down
```