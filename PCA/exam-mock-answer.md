# PCA (Prometheus Certified Associate) Exam Mock

## PCA
- https://training.linuxfoundation.org/certification/prometheus-certified-associate/

## Exam Info
- Duration: 90 min
- Question: 60
- Pass: x > 75% (min 45/60)

## Exam Topics

### Observability Concepts (18%)
- Prometheus is a pull-based time-series database
- Observability: understand what's happening inside a system and predict how it will behave in the future
- Monitoring: continues observation of a system to detect and alert on abnormal behavior.
- Telemetry: automates collection and transmission of data from remote source.
- RED Method consists of: (Request) Rate + (Request) Errors + (Request) Duration
- SLO: Service Level Objective (Goal), SLA: Service Level Agreement (Contract), SLI: Service Level Indicator (Metrics)
- Span is a single operation or unit of work within a distributed system and captures the start and end times, duration, and associated metadata of a specific operation
- 3 core components of observability are Logging, Trace and Metrics
- Metrics (numeric value) are monitorized by Prometheus
- Service Level Indicator (SLI) is typically derived from metrics
- An error budget policy is a concept commonly used in the context of SLOs and SLAs. The purpose of an error budget policy is to define the acceptable level of errors or service disruptions that a system or service can experience within a given time period.
- the advantage of push-system is that the data collection is timely and proactive (real-time or near real-time)
- InfluxDB is a pull-based time-series database designed to handle high volumes of time-stamped data (IoT, Sensor, Analytics).
- ELK stack is a push-based system, used for collecting, processing, storing, and visualizing log data.
- Prometheus is a pull-based time-series database and monitoring system specifically designed for monitoring dynamic cloud-native environments.
- numeric time-series data point 
- An exemplar is a specific trace representative of measurement taken in a given time interval.

### Prometheus Fundamentals (20%)
- promtool
- Prometheus has the following limitations: 1. scalability for large-scale deployments with millions of TS, 2. Long-term storage, 3. High cardinality, 4. HA and Replication
- SD is a mechanism that allow to automatically discover and monitor targets and services. There are 2 categories: top-down and bottom-up mechanisms of static SD
- scrape_interval
- scrape_configs -> relabel_configs -> action: drop
- with the flag `--storage.tsdb.retention.time` and `--storage.tsdb.retention.size`
- Retrieval, TSDB, HTTP Server
- 1.Sending a SIGHUP signal to the Prometheus process, 2.Using the Prometheus API `/-/reload`, 3. Using a service manager (systemctl) or orchestration tool (k8s)
- GET method
- `ec2_sd_configs`
- how frequently Prometheus collects and updates the metrics
- time-series database
- Prometheus exporter
- Pushgateway
- The 'honor_labels' configuration option gives you control over how labels are handled during metric ingestion.
- 9090:prometheus-server, 9093:altermanger, 9100:node-cluster, 9091:pushgateway
- instance,job
- ext4, XFS, and NTFS
- icmp -> prober:icmp
- scrape_configs > static_configs -> targets:xxx
- agent mode is a light promtheus mode, which is focused for remote-write (remote storage), service-discovery and scraping specially for edge-computing, iot.
- `./promtool test rules test.yml`
- `./promtool check rules test.yml`


### PromQL & Metrics (28%)
- Query Language for Prometheus
- Histogramm samples observations (e.g. request durations or response sizes) and counts them in configurable buckets
- Scalar, String, Instant Vector, Range Vector are used in PromQL
- Instant Vector
- `predict-linear()`
- `boolean`
- `avg_over_time()`
- `avg_over_time()` returns range vector and `avg()` returns aggregated num. avg_over_time has range vector as input and avg has instant vector
- `rate()` needs counter type metrics / range vector
- `offset` refers to the past time as duration
- `rate()` calc avg rate of change of a time series over the specified time range, irate() calc rate of change of a time series at each data point within the time range
- `deriv()` function calculates the instantaneous rate of change at each data point within the range.
- guage
- The data type of Prometheus metric values is float64.
- `_bucket`, `_sum` and `_count`
- metric name, label, timestamp, value
- floor=round a number down, ceil=round a number up


### Instrumentation and Exporters (16%)
- What is the HTTP headers to establish by Prometheus during each scrape?
- target + module
- Promtheus is using text-based format for exposing metrics
- No
- /metrics
- `build_info`
- Blackbox Exporter
- SNMP exporter
- Registry serves as a central repository for collecting, storing, and managing metrics
- Exporter is responsible for collecting metrics from a specific system, application, or service and exposing them for Prometheus
- Network Service Monitoring, Helth Check, Externe Monitoring
- scrape_configs > `metrics_path` et
- Blackbox Exporter allows blackbox probing of endpoints over HTTP, HTTPS, DNS, TCP, ICMP and gRPC
- HELP, TYPE


### Recording & Alerting & Dashboarding (18%)
- attribute in the route `time_intervals` ex. `time_intervals: [holidays, offhours]`. `mute_time_interval` is DEPRECATED.
- As a best practice, alerting should be triggered on symptoms and NOT causes
- The "what’s broken" indicates the symptom; the "why" indicates a (possibly intermediate) cause.
- symptom is customer visible error
- `<<level>>:<<metric>>:<<operations>>`, e.g.`job:node_cpu_seconds:avg_idle`
- acknowledge-based = notifications for an alert are sent to the recipient only once until the alert is acknowledged or resolved
- time-based = timiting the rate of notifications based on a specific time interval
- firing,pending,active,inactive
- not possible
- aggregate, transform, or filter metrics with PromQL and storing them into Prometheus DB
- Alert fatigue refers to a situation where individuals or teams become overwhelmed or desensitized by a large volume of alerts
- notification templates
- `group_by`
- grafana
- Inhibiting refers to a feature that allows certain alerts to be stopped or prevented from generating notifications for a specified duration of time
- YAML
- `for` allows for a delay or threshold before an alert is triggered, helping to prevent false positives and reduce noise in alerting systems
- Slience
- Prometheus Console
- Grouping
- Routing
- continue: specifies whether to continue processing subsequent routes after sending a notification for an alert.
- group_wait: sets how long to initially wait to send a notification
- repeat_interval: is used to determine the wait time before a firing alert that has already been successfully
- group_interval: dictates how long to wait before sending notifications about new alerts
- annotations, labels
- exemplar is a special label added to a TS that provides additional information about a specific data point.
- This can be configured using the `--cluster-*` flags