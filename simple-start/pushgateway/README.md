# Pushgateway

```bash
# run prometheus
docker-compose up -d

# run python app to push message
python app.py

# open prometheus and check batch job
open http://localhost:9091/
```