# Prometheus Golang Instrumentator

## Create a docker image and push into Dockerhub
```bash
export REPO=yuyatinnefeld
export IMAGE_NAME=go-prometheus-app:1.0.0

docker build -t $REPO/$IMAGE_NAME .
docker run -p 8888:8888 -d $REPO/$IMAGE_NAME

curl http://127.0.0.1:8888/
curl http://localhost:8888/metrics

docker push $REPO/$IMAGE_NAME

# test repo image
docker run -p 8888:8888 --name golang-prom $REPO/$IMAGE_NAME
curl http://127.0.0.1:8888
curl http://127.0.0.1:8888/metrics/
```

## Run Golang App via Docker-Compose
```bash
# run
docker-compose up -d

# verify
open http://localhost:8888/metrics

# verify
open http://localhost:9090/targets?search=my-golang-app

# verify
open http://localhost:9090/graph?g0.expr=go_app_requests_count
```
