# PCA (Prometheus Certified Associate) Exam Mock

## PCA
- https://training.linuxfoundation.org/certification/prometheus-certified-associate/

## Exam Info
- Duration: 90 min
- Question: 60
- Pass: x > 75% (min 45/60)

## Exam Topics

### Observability Concepts (18%)
- Observability: understand what's happening inside a system and predict how it will behave in the future
- Monitoring: continues observation of a system to detect and alert on abnormal behavior.
- Telemetry: automates collection and transmission of data from remote source.
- RED Method consists of: (Request) Rate + (Request) Errors + (Request) Duration

### Prometheus Fundamentals (20%)
- Prometheus has the following limitations: 1. scalability for large-scale deployments with millions of TS, 2. Long-term storage, 3. High cardinality, 4. HA and Replication
- InfluxDB is a pull-based time-series database designed to handle high volumes of time-stamped data (IoT, Sensor, Analytics).
- ELK stack is a push-based system, used for collecting, processing, storing, and visualizing log data.
- Prometheus is a pull-based time-series database and monitoring system specifically designed for monitoring dynamic cloud-native environments.
- SD is a mechanism that allow to automatically discover and monitor targets and services. There are 2 categories: top-down and bottom-up mechanisms of static SD


### PromQL & Metrics (28%)
- PromQL is read-only and used for aggregation of metrics
- The data types: string scalars, range vectors, and instant vectors are used in PromQL
- Histogramm samples observations (e.g. request durations or response sizes) and counts them in configurable buckets

### Instrumentation and Exporters (16%)
- Promtheus is using text-based format for exposing metrics
- Backbox Exporter allows blackbox probing of endpoints over HTTP, HTTPS, DNS, TCP, ICMP and gRPC

### Recording & Alerting & Dashboarding (18%)
- the naming convention for the recoring rules is: `<<level>>:<<metric>>:<<operations>>`, e.g.`job:node_cpu_seconds:avg_idle`
- As a best practice, alerting should be triggered on symptoms and NOT causes