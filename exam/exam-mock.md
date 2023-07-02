# Prometheus Certified Associate Practice Exam


## Exam Info
- Duration: 90 min
- Questions
- Question: 60
- Pass: x > 75% (min 45/60)

## Exam Topics

### Observability Concepts (18%)
- What are the distinctions between SLO, SLI, and SLA?
- In the context of tracing, what is the meaning or representation of a span?
- In which scenarios is distributed tracing less beneficial or not as applicable?
- What is typically tracked within a span of a trace?
- What are the 3 core components of observability?
- Which type of data is monitored by Prometheus?
- What is an error budget?
- In the context of monitoring and observability, what type of data is typically used to define a Service Level Indicator (SLI)?


### Prometheus Fundamentals (20%)
- Which property configures the timing for Prometheus to scrape new metrics from a target?
- What is the preferred approach used by Prometheus to collect metrics from a target?
- Which section in the Prometheus configuration file governs the selection of targets to be scraped?
- Which action in the label configuration is used to delete a specific target?
- How is managed data retention in prometheus?
- What are the essential components of Prometheus?
- Which two query parameters are required when configuring a Blackbox Exporter probe?
- What is a recognized method to initiate a configuration reload in Prometheus?
- What HTTP method does Prometheus employ for performing scrapes?
- Which service discovery configuration is recommended for scraping EC2 instances?
- What is the significance and purpose of the scrape_interval configuration in Prometheus?
- What does the term "inhibiting" refer to in the context of Prometheus?
- Which type of database does Prometheus utilize?
- What component is responsible for collecting metrics from an instance and exposing them in a format that Prometheus expects?
- Which component is suitable for collecting metrics from batch jobs?
- When is the configuration option 'honor_labels' used?
- What is the purpose of port 9090/9093/9100/9091 in Prometheus?
- what are 2 default metric labels?
- Which of the file systems is recommended/supported by Prometheus?
- How can you configure a Blackbox Exporter probe to check the successful response of your servers to ping?



### PromQL & Metrics (28%)
- What is the name of the vector in Prometheus that stores a single sample value for each time series, all of which have synchronized timestamps?
- Which PromQL function is used to estimate the value of a time series at a future time, t seconds from the current time, based on the range vector v?
- Between what type of expressions can logical operators be defined?
- Which function can be used to calculate the average of a range vector in Prometheus?
- With which type of metrics is the "rate()" function primarily used in Prometheus?
- What does the term "offset" refer to in Prometheus?
- What distinguishes the rate and irate query functions in Prometheus?
- What  is the definition of a metric?
- What is the CLI utility tool for Prometheus called?
- Which type of metric is suitable for measuring the internal temperature of a server?
- What is the data type of Prometheus metric values?




### Instrumentation and Exporters (16%)
- Does Prometheus need to perform any format conversion on the metrics returned by a monitored Linux machine?
- What is the default endpoint that Prometheus uses to scrape the metrics from the target?
- Where is the version of the Prometheus exporter typically defined?
- Which Prometheus exporter is recommended for monitoring network devices?
- What is the purpose of a Prometheus metrics registry?
- What is the purpose or definition of a Prometheus exporter?
- In what scenarios would you use the Blackbox Exporter?


### Recording & Alerting & Dashboarding (18%)
- What is Acknowledge-based throttling and Time-based throttling?
- What are the 3 possible statuses of a Prometheus alert?
- How can I use a PromQL query to retrieve the currently active alerts in Alertmanager?
- What is the meaning of "alert symptoms" and "alert causes" in the context of monitoring systems like Prometheus?
- What is the function of recording rules in Prometheus?
- Whas is the alert fatigue?
- Which feature of Alertmanager is responsible for formatting and customizing the alerts?
- How can you configure Alertmanager to disable the grouping of alerts for a specific route effectively?
- Which software is commonly used for visualizing Prometheus metrics?
- What is the function of the 'inhibiting' feature in Alertmanager?
- What is the format used for defining alerting rules?
- What is the significance of the "for" attribute in a Prometheus alert rule?
- How can you temporarily mute/snooze an alert during maintenance in Prometheus?
- What is the name of Prometheus' native dashboarding and visualization feature?
- How can you coordinate the simultaneous sending of multiple alerts with similar label sets in Prometheus?
- How can Alertmanager temporarily suppress notifications for specific alerts?
- Which feature of Alertmanager is resonsilbe for sending alert to the right receiver?










