# Prometheus

## Reference
- https://training.linuxfoundation.org/certification/prometheus-certified-associate/
- https://training.promlabs.com/

## Domains & Competencies

- Observability Concepts 18%
    - Metrics
    - Understand logs and events
    - Tracing and Spans
    - Push vs Pull
    - Service Discovery
    - Basics of SLOs, SLAs, and SLIs

- Prometheus Fundamentals 20%
    - System Architecture
    - Configuration and Scraping
    - Understanding Prometheus Limitations
    - Data Model and Labels
    - Exposition Format

- PromQL 28%
    - Selecting Data
    - Rates and Derivatives
    - Aggregating over time
    - Aggregating over dimensions
    - Binary operators
    - Histograms
    - Timestamp Metrics

- Instrumentation and Exporters 16%
    - Client Libraries
    - Instrumentation
    - Exporters
    - Structuring and naming metrics

- Alerting & Dashboarding 18%
    - Dashboarding basics
    - Configuring Alerting rules
    - Understand and Use Alertmanager
    - Alerting basics (when, what, and why)

## Prometheus vs InfluxDB vs ELK
- Prometheus = Time Series DBMS / Pull-based system / PROM <- App / used for Monitoring (metrics)
- InfluxDB = Time Series DB / Push-based system / App -> IFDB / used for IoT, Sensors, or Analytics Monitoring
- ELK = Search Engine / Logstash is rather Push-based system / used for Logging (logs)

## Main Components and Basic Terminologies
- Data Retrieval Worker is used for pulling metrics
- Time Series Database is for storing metrics
- HTTP Server is used for accepting PromQL queries and connecting to visualizataion tools like Grafana
- Alert/Alerting is outcome of an alerting rule which is firing from Prometheus to Alertmanager
- Alertmanager manages and sends out notifications to receiver like email, slack
- Target is a monitoring object (PostgreSQL, Nodejs, NodeExporter, Golang-App, etc.)
- Instance is an endpoint for scraping. localhost:8080/metrics
- Job is a collection of instances with the same purpose

## Observability vs Monitoring vs Telemetry
- Observability: understand what's happening inside a system and predict how it will behave in the future
- Monitoring: continues observation of a system to detect and alert on abnormal behavior.
- Telemetry: automates collection and transmission of data from remote source.
- Telemetry data: metrics, logs, traces
- Prometheus is a tool for gathering metrics

## PromQL
- Prometheus Querying Language
- read-only and is used for aggregation of metrics
- PromQL uses three data types: scalars, range vectors, and instant vectors.

## RED Method
- (Request) Rate - the number of requests, per second, you services are serving.
- (Request) Errors - the number of failed requests per second.
- (Request) Duration - distributions of the amount of time each request takes.

## Metric Types
- Counter: is cumulative metrics (going up always) and resets to zero though restart of server. you can use a counter to represent the number of requests served, tasks completed, or errors.
- Gauge: is an absolute metric that represents a single numerical value and can go up and down. e.g. current memory, temperature or pressure
- Histogramms: samples observations (e.g. request durations or response sizes) and counts them in configurable buckets
- Summaries: summary samples observations (e.g. request durations and response sizes)

## Instrumentation
- Service Instrumentation
    - Online Serving Systems
        - metrics (request rate, latency, error rate, in-progress-requests)
    - Offline Processing Systems (no one wating for a response)
        - metrics (progress count, errors, last exec time)
    - Batch Jobs (like Offline Processing but scheduled)
        - metrics (progress count, errors, last exec time)
- Library Instrumentation
    - metrics (internal erros, latency time withing the lib itself)

## Recording & Alert Rules
- You can create the recording rules though YAML files to manage frequently used complex PromQL expressions as new set of time series
- There are two types of rules: alerting rules and recording rules
- Benefits is the fast execution time. very helpful for dashboards
- Naming Recoring Rules: `<<level>>:<<metric>>:<<operations>>`, e.g.`job:node_cpu_seconds:avg_idle`
- Silences are a straightforward way to simply mute alerts for a given time
- Inhibition is a concept of suppressing notifications for certain alerts if certain other alerts are already firing.

## Backbox Exporter
- allows blackbox probing of endpoints over HTTP, HTTPS, DNS, TCP, ICMP and gRPC.

## Service Discovery
- SD is a mechanism that allow to automatically discover and monitor targets and services
- there are 2 categories: top-down and bottom-up mechanisms of static SD

## Service Monitors
- Service Monitors define a set of targets for prometheus to monitor and scrape
- Service Monitors allow you to avoid edditing prometheus configs directly and give you a DECLARATIVE k8s syntax to define targets

































