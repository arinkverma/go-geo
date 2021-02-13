# go-geo
Golang and redis based, high-performance, light weight micro-service to resolve location coordinates into cities. The source is based on **GeoNames**.

![Structure flow](https://github.com/arinkverma/go-geo/raw/main/high-perf-geo-service.png)

## Installation
```
make all

# Download data
make download

# Build
make build

# Run
make run
```

## Data source
The source is based on **GeoNames**. The GeoNames geographical database covers all countries and contains over eleven million placenames that are available for download free of charge.

## APIs

#### /resolve/{country-code}/{latitude},{longitude}

`curl http://0.0.0.0:8080/resolve/IN/28.612912,77.227321`
```{
	"city": "New Delhi",
	"country": "IN",
	"geonameId": "1261481"
}
```

`curl http://0.0.0.0:8080/resolve/IN/19.0821978,72.7410988`
```{
	"city": "Mumbai",
	"country": "IN",
	"id": "1275339"
}
```

`curl http://0.0.0.0:8080/resolve/CN/40.4319077,116.5681862`
```{
	"city": "Shunyi",
	"country": "CN",
	"geonameId": "2034754"
}
```

`curl http://0.0.0.0:8080/resolve/SG/1.33496,103.6961638`
```{
	"city": "Woodlands",
	"country": "SG",
	"geonameId": "1882316"
}
```

#### /ping
`curl http://0.0.0.0:8080/ping`
```{
	"message": "PONG"
}
```

## Redis geo engine
### GEOADD
* Time complexity: O(log(N)) for each item added, where N is the number of elements in the sorted set.
* There are limits to the coordinates that can be indexed: areas very near to the poles are not indexable.
* Example: 
```
redis> GEOADD Sicily 13.361389 38.115556 "Palermo" 15.087269 37.502669 "Catania"
(integer) 2
```

### GEORADIUS
* Time complexity: O(N+log(M)) where N is the number of elements inside the bounding box of the circular area delimited by center and radius and M is the number of items inside the index.
* Example:
```
redis> GEORADIUS Sicily 15 37 200 km
1) "Palermo"
2) "Catania"
```

## Performance on local VM with 2core
CPU: Intel(R) Core(TM) i5-6200U CPU @ 2.30GHz

Time per request: 2.896ms
- mean, across all concurrent requests

```
Concurrency Level:      50
Time taken for tests:   2.896 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      171000 bytes
HTML transferred:       48000 bytes
Requests per second:    345.36 [#/sec] (mean)
Time per request:       144.776 [ms] (mean)
Time per request:       2.896 [ms] (mean, across all concurrent requests)
Transfer rate:          57.67 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.6      0       4
Processing:    11  142  1re    143     173
Waiting:        8  141  15.5    141     173
Total:         11  142  15.3    143     174

Percentage of the requests served within a certain time (ms)
  50%    143
  66%    148
  75%    151
  80%    154
  90%    160
  95%    165
  98%    169
  99%    172
 100%    174 (longest request)
 ```
