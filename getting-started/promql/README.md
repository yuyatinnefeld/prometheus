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

## Hands-On Labs

### Instant vector and Range Vector
`./promql/PromQL-Lab-1.md`

### average
`./promql/PromQL-Lab-2.md`

### avg() vs avg_over_time()
`./promql/PromQL-Lab-3.md`

### Rate()
`./promql/PromQL-Lab-4.md`

### Histogram
`./promql/PromQL-Lab-5.md`

### Summary
`./promql/PromQL-Lab-6.md`


## PromQL Operator

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

