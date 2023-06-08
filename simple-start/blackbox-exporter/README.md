# Blackbox Exporter

Blackbox Exporter allows blackbox probing of endpoints over HTTP, HTTPS, DNS, TCP, ICMP and gRPC.


## Blackbox Config files
- https://github.com/prometheus/blackbox_exporter/releases/tag/v0.24.0

## Setup
```bash
# run prometheus + blackbox-expoter
docker-compose up -d

# verify blackbox expoter
open http://localhost:9115/

# check google.com
http://localhost:9115/probe?target=google.com&module=icmp_test&debug=true

# check yuyatinnefeld.com
http://localhost:9115/probe?target=yuyatinnefeld.com&module=icmp_test&debug=true

# check yuyatinnefeld.com dns
http://localhost:9115/probe?target=8.8.8.8&module=dns_test&debug=true

# check if blackbox exporter can scrape python-app
open http://localhost:8080/

# open blackbox exporter
http://localhost:9115/probe?target=localhost:8080&module=icmp_test&debug=true


open http://localhost:9115/probe?target=localhost:8080&module=tcp_connect

# Responsible for the scraping is: L19-22 | blackbox_config.yaml
#  tcp_connect:
#    prober: tcp
#    timeout: 5s


```