# PCA (Prometheus Certified Associate) Exam Mock

## PCA
- https://training.linuxfoundation.org/certification/prometheus-certified-associate/

## Exam Info
- Duration: 90 min
- Question: 60
- Pass: x > 75% (min 45/60)

## Exam Topics

### Observability Concepts (18%)
1. pull-based
2. Observability: understand what's happening inside a system and predict how it will behave in the future
2. Monitoring: continues observation of a system to detect and alert on abnormal behavior.
2. Telemetry: automates collection and transmission of data from remote source.
3. RED Method consists of: (Request) Rate + (Request) Errors + (Request) Duration
4. SLO: Service Level Objective (Goal), SLA: Service Level Agreement (Contract), SLI: Service Level Indicator (Metrics)
5. Span is a single operation/unit of work within a distributed system and captures the start and end times, duration, and associated metadata of a specific operation
6. for monolith system
7. Operation Name, Trace ID and Span ID, Start and End Timestamps, Duration, Parent Span ID
8. Logging, Trace and Metrics
9. Metrics (numeric value)
10. SLI is typically derived from metrics
11. An error budget policy is a concept used in the context of SLOs and SLAs and is to define the acceptable level of errors or service disruptions that a system or service can experience within a given time period.
12. timely and proactive data collection (real-time or near real-time) / pushing into the centralized data system
13. InfluxDB is a pull-based time-series database designed to handle high volumes of time-stamped data (IoT, Sensor, Analytics).
13. ELK stack is a push-based system, used for collecting, processing, storing, and visualizing log data.
13. Prometheus is a pull-based time-series database and monitoring system specifically designed for monitoring dynamic cloud-native environments.
14. numeric time-series data point
15. An exemplar is a specific trace representative of measurement taken in a given time interval and provides additional information about a specific data point.
16. To gather and aggregate textual event data from a service for troubleshooting

### Prometheus Fundamentals (20%)
1. promtool
2. scalability for large-scale deployments with millions of TS, Long-term storage, High cardinality, HA and Replication
3. SD is a mechanism that allow to automatically discover and monitor targets and services. There are 2 categories: top-down (e.i. ec2) and bottom-up (e.i. consol) mechanisms of static SD
4. `scrape_interval`
5. `scrape_configs`
6. `scrape_configs` -> `relabel_configs` -> `action: drop` or `action: keep`
7. with the flag `--storage.tsdb.retention.time` and `--storage.tsdb.retention.size`
8. Retrieval, TSDB, HTTP Server
9. with the flag `--web.enable-lifecycle`
10. Sending a SIGHUP signal to the Prometheus process, Using the Prometheus API `POST or PUT + /-/reload`, Using a service manager (systemctl) or orchestration tool (k8s)
11. HTTP GET method
12. `ec2_sd_configs`
13. `ec2_sd_configs`
14. how frequently Prometheus collects and updates the metrics
15. time-series database
16. Prometheus exporter
17. Pushgateway
18. Using `honor_labels` can make your collected metrics more informative and allow you to differentiate between different metrics coming from various sources or probe targets
19. 9090:prometheus-server, 9093:altermanger, 9100:node-exporter, 9091:pushgateway, 9115:blackbox-exporter
20. `instance` and `job`
21. ext4, XFS, and NTFS
22. Internet Control Message Protocol (ICMP) -> `prober:icmp`
23. scrape_configs > static_configs -> targets:xxx
24. agent mode is a light promtheus mode, which is focused for remote-write (remote storage), service-discovery and scraping specially for edge-computing/IoT and reducing for querying, alerting and local storage 
25. `./promtool test rules test.yml`
26. `./promtool check rules test.yml`
27. `scrape_configs` and `*_sd_configs` on per-job basis
28. starting the server with the flag `--web.enable-admin-api` + `curl - X POST -g 'http://localhost:9090/api/v1/admin/tsdb/delete_series?match[]={xxxx="yyy"}'`
29. starting the server with the flag `--web.enable-admin-api` + `$ curl -X POST -g 'http://localhost:9090/api/v1/admin/tsdb/clean_tombstones'`
30. YAML and JSON

### PromQL & Metrics (28%)
1. Query Language for Prometheus
2. Histogramm samples observations (e.g. request durations or response sizes) and counts them in configurable buckets
3. Scalar, String, Instant Vector, Range Vector
4. Instant Vector
5. `predict-linear()`
6. `boolean`
7. `avg_over_time(metrics[x])`
8. `avg_over_time(...)` has range vector as input and returns range vector as output. `avg(...)` has instant vector as input and returns aggregated number.
9. `rate(...)` needs COUNTER type metrics
10. `offset` refers to the past time as duration
11. `rate(...)` calc avg rate of change of a time series over the specified time range, `irate(...)` calc avg rate of change of a time series at the last 2 data points
12. `deriv(...)` operates on gauge and `rate(...)` operates on counter
13. guage
14. float64.
15. `<basename>_bucket`, `<basename>_sum` and `<basename>_count`
16. metric name, metrics label, timestamp, value
17. `floor(...)` = round a number down, `ceil(...)` = round a number up
18. `absent(...)`
19. logical => `OR, AND, UNLESS`, arithmetic => `+ - * / % ^`, comparison => `==, !=, >, <, >=, <=`
20. `on, ignoring`
21. a part of vector matching. `on, ignoring` + `group_left, group_right`
22. `idelta()`
23. `max(cert_expiry - time()) / 86400`
24. `sum(), min(), max(), avg(), count()`
25. The label is a reserved label

### Instrumentation and Exporters (16%)
1. `X-Prometheus-Scrape-Timeout-Seconds`
2. `target` + `module`
3. text-based format for exposing metrics
4. No
5. `/metrics`
6. `build_info`
7. Blackbox Exporter
8. SNMP exporter
9. HTTP protocol 
10. Registry serves as a central repository for collecting, storing, and managing metrics
11. Exporter is responsible for collecting metrics from a specific system, application, or service and exposing them for Prometheus
12. Network Service Monitoring, Helth Check, Externe Monitoring
13. `scrape_configs` > `metrics_path: /metrics`
14. Blackbox Exporter allows blackbox probing of endpoints over `HTTP, HTTPS, DNS, TCP, ICMP, gRPC`
15. `file_sd_configs`
16. HELP, TYPE
17. JMX Exporter
18. `honor_labels:true`
19. PromQL > `job_last_success_unixtime`

### Recording & Alerting & Dashboarding (18%)
1. attribute in the route `time_intervals` ex. `time_intervals: [holidays, offhours]`. `mute_time_interval` is DEPRECATED.
2. symptom-based and NOT causes-based
3. sympton: The "what’s broken", cause: "why broken"
4. symptom is customer visible error
5. `<<level>>:<<metric>>:<<operations>>`, e.g.`job:node_cpu_seconds:avg_idle`
6. acknowledge-based = notifications for an alert are sent to the recipient only once until the alert is acknowledged or resolved
6. time-based = timiting the rate of notifications based on a specific time interval (ex. goup_interval, scrape_interval)
7. firing, pending, inactive
8. Mene > Alerts > Query > `ALERTS`
9. aggregate and filter metrics with PromQL and storing them into Prometheus DB
10. `rules` -> `record: xxx`,`expr: xxx`
11. Alert fatigue refers to a situation where individuals or teams become overwhelmed or desensitized by a large volume of alerts
12. notification templates
13. `group_by`
14. Grafana
15. Inhibiting refers to a feature that allows certain alerts to be stopped or prevented from generating notifications for a specified duration of time
16. YAML
17. `for` allows for a delay or threshold before an alert is firing, helping to prevent false positives and reduce noise in alerting systems
18. Slience
19. Prometheus Console
20. Grouping
21. Routing
22. `repeat_interval`: is used to determine the wait time before a firing alert that has already been successfully
22. `continue`: specifies whether to continue processing subsequent routes after sending a notification for an alert
22. `group_wait`: sets how long to initially wait to send a notification
22. `group_interval`: dictates how long to wait before sending notifications about new alerts
23. `annotations + labels`
24. This can be configured using the `--cluster-*` flags
25. Active, Pending, Expired