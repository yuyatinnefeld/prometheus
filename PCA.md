# Prometheus

## Reference
- https://training.linuxfoundation.org/certification/prometheus-certified-associate/
- https://training.promlabs.com/

## Domains & Competencies
- Observability Concepts 18%
- Prometheus Fundamentals 20%
- PromQL 28%
- Instrumentation and Exporters 16%
- Alerting & Dashboarding 18%

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
- Gauge: is an absolute metric that represents a single numerical value and can go up and down. e.g. current memory or cpu usage
- Histogramms: samples observations (e.g. request durations or response sizes) and counts them in configurable buckets
- Summaries: summary samples observations (e.g. request durations and response sizes)

## Service Monitors
- Service Monitors define a set of targets for prometheus to monitor and scrape
- Service Monitors allow you to avoid edditing prometheus configs directly and give you a DECLARATIVE k8s syntax to define targets


## Exam Dumps
- Which property sets the time after which Prometheus will scrape the new metrics from target?
- Which approach Prometheus prefers to scrape metrics from a target?
- Which block in the Prometheus configuration file controls what targets are to be scraped?
- Suppose you have a Linux machine being monitored. Does Prometheus has to convert the format of metrics being returned.
- From which endpoint Prometheus scrapes the target metrics by default?


















