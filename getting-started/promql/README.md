# PromQL

- Cheat Sheet: https://promlabs.com/promql-cheat-sheet/
- Examples: https://prometheus.io/docs/prometheus/latest/querying/examples/

## Metric Types
- counters
- gauges
- histograms
- summaries

## Data Types
- Range vectors
- Instant vectors
- Scalar
- String

### Instant vector
- a set of time series containing a SINGLE SAMPLE

Exp: `<<metric name>>{<<label>>=”<<value>>”}` like `promhttp_metric_handler_requests_total{code="200",job="node"}`
```bash
# use demo prometheus ui site
curl https://demo.promlabs.com/metrics

# Instant Vector:
curl 'https://demo.promlabs.com/api/v1/query' \
    --data 'query=promhttp_metric_handler_requests_total{code="200",job="node"}' | jq
```

### Range vector
- a set of time series containing a RANGE OF DATA POINTS
Exp: `<<metric name>>{<<label>>=”<<value>>”}[duration]` like `promhttp_metric_handler_requests_total{code="200",job="node"}[30s]`

```bash
# Range Vector:
curl 'https://demo.promlabs.com/api/v1/query' \
    --data 'query=promhttp_metric_handler_requests_total{code="200",job="node"}[30s]' | jq
```

### Scalar
- a simple numeric floating point value

### String
- a simple string value; currently unused


## PromQL Operators

### Selectors & Matchers
```bash
# equality matcher (=)
curl 'https://demo.promlabs.com/api/v1/query' \
    --data 'query=process_cpu_seconds_total{job="node"}' | jq

#negative equality matcher (!=)
curl 'https://demo.promlabs.com/api/v1/query' \
    --data 'query=process_cpu_seconds_total{job!="node"}' | jq

# regex matcher (=~)
curl 'https://demo.promlabs.com/api/v1/query' \
    --data 'query=process_cpu_seconds_total{job=~".*prom.*"}' | jq

# negative regex matcher (!~)
curl 'https://demo.promlabs.com/api/v1/query' \
    --data 'query=process_cpu_seconds_total{job!~".*prom.*"}' | jq
```

### Arithmetic Binary Operators

```bash
+ (add)
– (subtract)
* (multiply)
/ (divide)
% (percentage)
^ (exponents)
```

### Comparison Binary Operators

```bash
== (equal to)
!= (does not equal)
> (greater than)
< (less than)
>= (greater than or equal to)
<= (less than or equal to)
```

### Logical/set Binary Operators
```bash
# intersection
and
# union
or
# complement
unless
```

### Aggregation Operators

```bash
sum 
avg
min 
max
group 
count 
count_values 
topk (k = the number of elements; this selects the largest values among those elements)
bottomk (like topk but for lowest values)
quantile (calculate a quantile over dimensions)
stddev (standard deviation over dimensions)
stdvar (standard variance over dimensions)

# examples
COUNT(prometheus_http_requests_total{code=~"2.*|4.*"}) BY (code)
MAX(prometheus_http_requests_total{code=~"4.*"}) BY (handler)
AVG(prometheus_http_requests_total) BY (code)
# 
```

### PromQL Functions
```bash
abs(instant-vector)
absent(instant-vector)
absent_over_time(range-vector)
ceil(instant-vector)
changes(range-vector)
clamp_max(instant-vector, scalar)
clamp_min(instant-vector, scalar)
day_of_month(some vector(time()) instant-vector)
day_of_week(some vector(time()) instant-vector)
days_in_month(some vector(time()) instant-vector)
delta(range-vector) #for use with gauge metrics
deriv(range-vector) #for use with gauge metrics
exp(instant-vector)
floor(instant-vector)
histogram_quantile(scalar, instant-vector)
holt_winters(range-vector, scalar, scalar)
hour(some vector(time()) instant-vector)
idelta(range-vector)
increase(range-vector)
irate(range-vector)
label_join()
label_replace()
ln(instant-vector)
log2(instant-vector)
log10(instant-vector)
minute(some vector(time()) instant-vector)
month(some vector(time()) instant-vector)
predict_linear() #for use with gauge 
rate(range-vector) # rate() only with range-vector
resets(range-vector) #for use with counter metrics
round((instant-vector, to_nearest=## scalar)
scalar(instant-vector)
sort(instant-vector)
sort_desc()
sqrt(instant-vector)
time()
timestamp(instant-vector)
vector()
year(some vector(time()) instant-vector)
avg_over_time(range-vector)
min_over_time(range-vector)
max_over_time(range-vector)
sum_over_time(range-vector)
count_over_time(range-vector)
quantile_over_time(scalar, range-vector) #φ-quantile (0 ≤ φ ≤ 1) of an interval’s values
stddev_over_time(range-vector) #standard deviation
stdvar_over_time(range-vector) #standard variance
```

### rate() vs irate()
- rate() caculates the `PER-SECOND-AVERAGE RATE` of increase of the time series in the range vector
- irate() cacluates the `INSTANT RATE` of increase of the time series in the range vector

- rate() is generally used when graphing the slow moving counters. While irate() is used when graphing the high volatile counters.

- irate uses only two last data points on each interval between subsequent pixels (aka step), while rate calculates average per-second rate over all the data points on each interval.



