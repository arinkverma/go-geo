# go-geo

## Installation
```
make all
```

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
```redis> GEOADD Sicily 13.361389 38.115556 "Palermo" 15.087269 37.502669 "Catania"
(integer) 2
```

### GEORADIUS
* Time complexity: O(N+log(M)) where N is the number of elements inside the bounding box of the circular area delimited by center and radius and M is the number of items inside the index.
* Example:
```redis> GEOADD Sicily 13.361389 38.115556 "Palermo" 15.087269 37.502669 "Catania"
(integer) 2
redis> GEORADIUS Sicily 15 37 200 km WITHDIST
1) 1) "Palermo"
   2) "190.4424"
2) 1) "Catania"
   2) "56.4413"
```
