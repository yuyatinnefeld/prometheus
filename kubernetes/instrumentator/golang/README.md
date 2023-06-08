# Prometheus Golang Instrumentator

## Init Setup
- we are using the Docker Image: yuyatinnefeld/go-prometheus-app:1.0.0
- this image is created and pushed: ./prometheus/instrumentator/golang-app

## Run Minikube
```bash
minikube start --cpus 4 --memory 8192
```

## Create Namespace
```bash
NS=prometheus
kubectl create namespace $NS
kubectl config set-context --current --namespace=$NS
```

## Deploy Golang Web App
```bash
kubectl apply -f golang.yaml -n $NS
kubectl get deploy
kubectl get svc

# test
kubectl port-forward svc/golang-app-svc 8888
curl http://127.0.0.1:8888/metrics/
```

## Run Prometheus
```bash
NS=prometheus
helm install prometheus-stack prometheus-community/kube-prometheus-stack -n $NS
kubectl get all -n $NS

# deploy service monitor for scraping
kubectl apply -f service-monitor.yaml -n $NS

# check service discovery and target
kubectl port-forward svc/prometheus-stack-kube-prom-prometheus 9090
```
