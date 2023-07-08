# Functions

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