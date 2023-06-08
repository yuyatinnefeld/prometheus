# Initial Setup

```bash
# start servers
docker-compose up -d

 # create metrics
open http://localhost:9100/metrics

# open prometheus
open http://localhost:9090/targets

# stop
docker-compose down
```