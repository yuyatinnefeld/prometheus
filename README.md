# Prometheus

## Architecture and Components
![Screenshot](pics/architecture.png)

## About
Prometheus collects and stores its log-metrics as time series data, i.e. metrics information is stored with the timestamp at which it was recorded, alongside optional key-value pairs called labels.

## Prep PCA Exam
- `/exma/PCA.md`
- `/exma/exam-mock.md`

## Hands-On Trainging

### Prometheus UI
- Demo Prometheus UI: https://demo.promlabs.com/
- Demo Proemtheus Metrics: https://demo.promlabs.com/metrics

### Getting Started
1. `/simple-start/setup`
2. `/simple-start/instrumentator`
3. `/simple-start/recording-rules`
4. `/simple-start/alert-rules`
5. `/simple-start/alertmanager`
6. `/simple-start/blackbox-exporter`
7. `/simple-start/pushgateway`
8. `/simple-start/service-discovery`
9. `/simple-start/promql`

#### Setup Prometheus with k8s
1. `/kubernetes/setup`
2. `/kubernetes/exporter`
3. `/kubernetes/instrumentator`

#### AWS
1. `/aws/deploy_prometheus_ec2_server`
2. `/aws/scrape_ec2_instance`

#### Setup Grafana
- `/kubernetes/grafana`

#### PromQL
- `/exam/promql.md`
