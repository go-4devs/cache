# cache

[![Build Status](https://drone.gitoa.ru/api/badges/go-4devs/cache/status.svg)](https://drone.gitoa.ru/go-4devs/cache)
[![Go Report Card](https://goreportcard.com/badge/gitoa.ru/go-4devs/cache)](https://goreportcard.com/report/gitoa.ru/go-4devs/cache)
[![GoDoc](https://godoc.org/gitoa.ru/go-4devs/cache?status.svg)](http://godoc.org/gitoa.ru/go-4devs/cache)

## Benchmark cache

`go test -v -timeout 25m -cpu 1,2,4,8,16 -benchmem -run=^$ -bench . bench_test.go`

```bash
goos: darwin
goarch: amd64
BenchmarkCacheGetStruct
BenchmarkCacheGetStruct/encoding_json
BenchmarkCacheGetStruct/encoding_json            	 1519932	       783 ns/op	     320 B/op	       6 allocs/op
BenchmarkCacheGetStruct/encoding_json-2          	 1478414	       780 ns/op	     320 B/op	       6 allocs/op
BenchmarkCacheGetStruct/encoding_json-4          	 1353025	       916 ns/op	     320 B/op	       6 allocs/op
BenchmarkCacheGetStruct/encoding_json-8          	 1284042	       839 ns/op	     320 B/op	       6 allocs/op
BenchmarkCacheGetStruct/encoding_json-16         	 1422788	       848 ns/op	     320 B/op	       6 allocs/op
BenchmarkCacheGetStruct/encoding_gob
BenchmarkCacheGetStruct/encoding_gob             	   83661	     15323 ns/op	    6944 B/op	     180 allocs/op
BenchmarkCacheGetStruct/encoding_gob-2           	   81745	     14407 ns/op	    6944 B/op	     180 allocs/op
BenchmarkCacheGetStruct/encoding_gob-4           	   73537	     15142 ns/op	    6944 B/op	     180 allocs/op
BenchmarkCacheGetStruct/encoding_gob-8           	   85412	     14494 ns/op	    6944 B/op	     180 allocs/op
BenchmarkCacheGetStruct/encoding_gob-16          	   75748	     15219 ns/op	    6944 B/op	     180 allocs/op
BenchmarkCacheGetStruct/map
BenchmarkCacheGetStruct/map                      	 6162325	       199 ns/op	      96 B/op	       1 allocs/op
BenchmarkCacheGetStruct/map-2                    	 5740689	       195 ns/op	      96 B/op	       1 allocs/op
BenchmarkCacheGetStruct/map-4                    	 6018531	       200 ns/op	      96 B/op	       1 allocs/op
BenchmarkCacheGetStruct/map-8                    	 5452492	       210 ns/op	      96 B/op	       1 allocs/op
BenchmarkCacheGetStruct/map-16                   	 5933622	       202 ns/op	      96 B/op	       1 allocs/op
BenchmarkCacheGetStruct/map_shards
BenchmarkCacheGetStruct/map_shards               	 5299807	       230 ns/op	      96 B/op	       1 allocs/op
BenchmarkCacheGetStruct/map_shards-2             	 5087726	       238 ns/op	      96 B/op	       1 allocs/op
BenchmarkCacheGetStruct/map_shards-4             	 4990490	       243 ns/op	      96 B/op	       1 allocs/op
BenchmarkCacheGetStruct/map_shards-8             	 4899127	       225 ns/op	      96 B/op	       1 allocs/op
BenchmarkCacheGetStruct/map_shards-16            	 5229320	       233 ns/op	      96 B/op	       1 allocs/op
BenchmarkCacheGetStruct/ristretto
BenchmarkCacheGetStruct/ristretto                	 5511872	       227 ns/op	      96 B/op	       1 allocs/op
BenchmarkCacheGetStruct/ristretto-2              	 4664298	       257 ns/op	     103 B/op	       1 allocs/op
BenchmarkCacheGetStruct/ristretto-4              	 4524751	       265 ns/op	     103 B/op	       1 allocs/op
BenchmarkCacheGetStruct/ristretto-8              	 4425381	       260 ns/op	     103 B/op	       1 allocs/op
BenchmarkCacheGetStruct/ristretto-16             	 4649698	       258 ns/op	     103 B/op	       1 allocs/op
BenchmarkCacheGetStruct/lru
BenchmarkCacheGetStruct/lru                      	 4730811	       250 ns/op	     144 B/op	       2 allocs/op
BenchmarkCacheGetStruct/lru-2                    	 4627194	       252 ns/op	     144 B/op	       2 allocs/op
BenchmarkCacheGetStruct/lru-4                    	 4627082	       257 ns/op	     144 B/op	       2 allocs/op
BenchmarkCacheGetStruct/lru-8                    	 4755622	       252 ns/op	     144 B/op	       2 allocs/op
BenchmarkCacheGetStruct/lru-16                   	 4717584	       250 ns/op	     144 B/op	       2 allocs/op
BenchmarkCacheGetStruct/redis_json
BenchmarkCacheGetStruct/redis_json               	     572	   2132479 ns/op	    9848 B/op	      34 allocs/op
BenchmarkCacheGetStruct/redis_json-2             	     565	   2161113 ns/op	    9848 B/op	      34 allocs/op
BenchmarkCacheGetStruct/redis_json-4             	     543	   2183219 ns/op	    9848 B/op	      34 allocs/op
BenchmarkCacheGetStruct/redis_json-8             	     531	   2148630 ns/op	    9848 B/op	      34 allocs/op
BenchmarkCacheGetStruct/redis_json-16            	     544	   2212659 ns/op	    9848 B/op	      34 allocs/op
BenchmarkCacheGetStruct/redis_gob
BenchmarkCacheGetStruct/redis_gob                	     553	   2206583 ns/op	   16504 B/op	     208 allocs/op
BenchmarkCacheGetStruct/redis_gob-2              	     549	   2256638 ns/op	   16505 B/op	     208 allocs/op
BenchmarkCacheGetStruct/redis_gob-4              	     540	   2230342 ns/op	   16504 B/op	     208 allocs/op
BenchmarkCacheGetStruct/redis_gob-8              	     537	   2178895 ns/op	   16504 B/op	     208 allocs/op
BenchmarkCacheGetStruct/redis_gob-16             	     541	   2206298 ns/op	   16504 B/op	     208 allocs/op
BenchmarkCacheGetStruct/memcache_json
BenchmarkCacheGetStruct/memcache_json            	    1352	    882575 ns/op	     560 B/op	      16 allocs/op
BenchmarkCacheGetStruct/memcache_json-2          	    1332	    869724 ns/op	     560 B/op	      16 allocs/op
BenchmarkCacheGetStruct/memcache_json-4          	    1326	    824555 ns/op	     561 B/op	      16 allocs/op
BenchmarkCacheGetStruct/memcache_json-8          	    1375	    880741 ns/op	     562 B/op	      16 allocs/op
BenchmarkCacheGetStruct/memcache_json-16         	    1346	    872861 ns/op	     563 B/op	      16 allocs/op
BenchmarkCacheGetStruct/memcache_gob
BenchmarkCacheGetStruct/memcache_gob             	    1431	    828348 ns/op	    7216 B/op	     190 allocs/op
BenchmarkCacheGetStruct/memcache_gob-2           	    1266	    875339 ns/op	    7216 B/op	     190 allocs/op
BenchmarkCacheGetStruct/memcache_gob-4           	    1327	    908142 ns/op	    7218 B/op	     190 allocs/op
BenchmarkCacheGetStruct/memcache_gob-8           	    1286	    840878 ns/op	    7219 B/op	     190 allocs/op
BenchmarkCacheGetStruct/memcache_gob-16          	    1540	    797765 ns/op	    7220 B/op	     190 allocs/op
```

## Benchmark providers

`go test -v -timeout 25m -cpu 1,2,4,8,16 -benchmem -run=^$ -bench . ./provider/bench_provider_test.go`

```bash
goos: darwin
goarch: amd64
BenchmarkCacheGetRandomKeyString
BenchmarkCacheGetRandomKeyString/encoding
BenchmarkCacheGetRandomKeyString/encoding            	 3100226	       389 ns/op	     192 B/op	       2 allocs/op
BenchmarkCacheGetRandomKeyString/encoding-2          	 3142849	       379 ns/op	     192 B/op	       2 allocs/op
BenchmarkCacheGetRandomKeyString/encoding-4          	 3118212	       379 ns/op	     192 B/op	       2 allocs/op
BenchmarkCacheGetRandomKeyString/encoding-8          	 3064170	       387 ns/op	     192 B/op	       2 allocs/op
BenchmarkCacheGetRandomKeyString/encoding-16         	 3128031	       384 ns/op	     192 B/op	       2 allocs/op
BenchmarkCacheGetRandomKeyString/map
BenchmarkCacheGetRandomKeyString/map                 	 7342993	       157 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetRandomKeyString/map-2               	 7268864	       158 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetRandomKeyString/map-4               	 7233045	       162 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetRandomKeyString/map-8               	 7393652	       159 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetRandomKeyString/map-16              	 7463053	       159 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetRandomKeyString/shard
BenchmarkCacheGetRandomKeyString/shard               	 3330136	       351 ns/op	      64 B/op	       1 allocs/op
BenchmarkCacheGetRandomKeyString/shard-2             	 3518775	       335 ns/op	      64 B/op	       1 allocs/op
BenchmarkCacheGetRandomKeyString/shard-4             	 3477537	       336 ns/op	      64 B/op	       1 allocs/op
BenchmarkCacheGetRandomKeyString/shard-8             	 3514064	       335 ns/op	      64 B/op	       1 allocs/op
BenchmarkCacheGetRandomKeyString/shard-16            	 3412119	       341 ns/op	      64 B/op	       1 allocs/op
BenchmarkCacheGetRandomKeyString/lru
BenchmarkCacheGetRandomKeyString/lru                 	 5013633	       249 ns/op	      48 B/op	       1 allocs/op
BenchmarkCacheGetRandomKeyString/lru-2               	 4871456	       247 ns/op	      48 B/op	       1 allocs/op
BenchmarkCacheGetRandomKeyString/lru-4               	 4786940	       238 ns/op	      48 B/op	       1 allocs/op
BenchmarkCacheGetRandomKeyString/lru-8               	 4721556	       238 ns/op	      48 B/op	       1 allocs/op
BenchmarkCacheGetRandomKeyString/lru-16              	 4870622	       241 ns/op	      48 B/op	       1 allocs/op
BenchmarkCacheGetRandomKeyString/ristretto
BenchmarkCacheGetRandomKeyString/ristretto           	 5569208	       205 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetRandomKeyString/ristretto-2         	 3892068	       295 ns/op	       7 B/op	       0 allocs/op
BenchmarkCacheGetRandomKeyString/ristretto-4         	 4490196	       266 ns/op	       7 B/op	       0 allocs/op
BenchmarkCacheGetRandomKeyString/ristretto-8         	 4381441	       266 ns/op	       7 B/op	       0 allocs/op
BenchmarkCacheGetRandomKeyString/ristretto-16        	 4185096	       273 ns/op	       7 B/op	       0 allocs/op
BenchmarkCacheGetRandomKeyString/memcache
BenchmarkCacheGetRandomKeyString/memcache            	    1492	    811587 ns/op	     528 B/op	      12 allocs/op
BenchmarkCacheGetRandomKeyString/memcache-2          	    1400	    840429 ns/op	     528 B/op	      12 allocs/op
BenchmarkCacheGetRandomKeyString/memcache-4          	    1381	    793654 ns/op	     528 B/op	      12 allocs/op
BenchmarkCacheGetRandomKeyString/memcache-8          	    1455	    826461 ns/op	     530 B/op	      12 allocs/op
BenchmarkCacheGetRandomKeyString/memcache-16         	    1380	    803712 ns/op	     532 B/op	      12 allocs/op
BenchmarkCacheGetRandomKeyString/redis
BenchmarkCacheGetRandomKeyString/redis               	     540	   2908289 ns/op	    9704 B/op	      30 allocs/op
BenchmarkCacheGetRandomKeyString/redis-2             	     514	   2287030 ns/op	    9704 B/op	      30 allocs/op
BenchmarkCacheGetRandomKeyString/redis-4             	     542	   2195917 ns/op	    9704 B/op	      30 allocs/op
BenchmarkCacheGetRandomKeyString/redis-8             	     536	   2209508 ns/op	    9704 B/op	      30 allocs/op
BenchmarkCacheGetRandomKeyString/redis-16            	     544	   2275867 ns/op	    9704 B/op	      30 allocs/op
BenchmarkCacheGetRandomKeyString/pebble
BenchmarkCacheGetRandomKeyString/pebble              	  672912	      1801 ns/op	    1408 B/op	       6 allocs/op
BenchmarkCacheGetRandomKeyString/pebble-2            	  773318	      1691 ns/op	    1408 B/op	       6 allocs/op
BenchmarkCacheGetRandomKeyString/pebble-4            	  729020	      1556 ns/op	    1408 B/op	       6 allocs/op
BenchmarkCacheGetRandomKeyString/pebble-8            	  778066	      1491 ns/op	    1408 B/op	       6 allocs/op
BenchmarkCacheGetRandomKeyString/pebble-16           	  838596	      1441 ns/op	    1408 B/op	       6 allocs/op
BenchmarkCacheGetRandomKeyInt
BenchmarkCacheGetRandomKeyInt/encoding
BenchmarkCacheGetRandomKeyInt/encoding               	 2825020	       410 ns/op	     207 B/op	       2 allocs/op
BenchmarkCacheGetRandomKeyInt/encoding-2             	 2932910	       409 ns/op	     207 B/op	       2 allocs/op
BenchmarkCacheGetRandomKeyInt/encoding-4             	 2837827	       408 ns/op	     207 B/op	       2 allocs/op
BenchmarkCacheGetRandomKeyInt/encoding-8             	 2842040	       418 ns/op	     207 B/op	       2 allocs/op
BenchmarkCacheGetRandomKeyInt/encoding-16            	 2866555	       409 ns/op	     207 B/op	       2 allocs/op
BenchmarkCacheGetRandomKeyInt/map
BenchmarkCacheGetRandomKeyInt/map                    	 7312549	       150 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetRandomKeyInt/map-2                  	 7884612	       150 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetRandomKeyInt/map-4                  	 7450554	       158 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetRandomKeyInt/map-8                  	 7471407	       156 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetRandomKeyInt/map-16                 	 7469587	       158 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetRandomKeyInt/shard
BenchmarkCacheGetRandomKeyInt/shard                  	 6709964	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetRandomKeyInt/shard-2                	 6430581	       183 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetRandomKeyInt/shard-4                	 6375858	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetRandomKeyInt/shard-8                	 6399346	       180 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetRandomKeyInt/shard-16               	 6580282	       175 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetRandomKeyInt/lru
BenchmarkCacheGetRandomKeyInt/lru                    	 5183596	       225 ns/op	      48 B/op	       1 allocs/op
BenchmarkCacheGetRandomKeyInt/lru-2                  	 5217847	       220 ns/op	      48 B/op	       1 allocs/op
BenchmarkCacheGetRandomKeyInt/lru-4                  	 5078146	       223 ns/op	      48 B/op	       1 allocs/op
BenchmarkCacheGetRandomKeyInt/lru-8                  	 4722044	       225 ns/op	      48 B/op	       1 allocs/op
BenchmarkCacheGetRandomKeyInt/lru-16                 	 4989286	       224 ns/op	      48 B/op	       1 allocs/op
BenchmarkCacheGetRandomKeyInt/ristretto
BenchmarkCacheGetRandomKeyInt/ristretto              	 6920838	       169 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetRandomKeyInt/ristretto-2            	 4763511	       216 ns/op	       7 B/op	       0 allocs/op
BenchmarkCacheGetRandomKeyInt/ristretto-4            	 5163074	       220 ns/op	       7 B/op	       0 allocs/op
BenchmarkCacheGetRandomKeyInt/ristretto-8            	 5133212	       220 ns/op	       7 B/op	       0 allocs/op
BenchmarkCacheGetRandomKeyInt/ristretto-16           	 5089780	       219 ns/op	       7 B/op	       0 allocs/op
BenchmarkCacheGetRandomKeyInt/memcache
BenchmarkCacheGetRandomKeyInt/memcache               	    1332	    820272 ns/op	     544 B/op	      14 allocs/op
BenchmarkCacheGetRandomKeyInt/memcache-2             	    1408	    840124 ns/op	     544 B/op	      14 allocs/op
BenchmarkCacheGetRandomKeyInt/memcache-4             	    1443	    809845 ns/op	     544 B/op	      14 allocs/op
BenchmarkCacheGetRandomKeyInt/memcache-8             	    1449	    832162 ns/op	     545 B/op	      14 allocs/op
BenchmarkCacheGetRandomKeyInt/memcache-16            	    1333	    855560 ns/op	     547 B/op	      14 allocs/op
BenchmarkCacheGetRandomKeyInt/redis
BenchmarkCacheGetRandomKeyInt/redis                  	     525	   2211523 ns/op	    9767 B/op	      31 allocs/op
BenchmarkCacheGetRandomKeyInt/redis-2                	     542	   2146253 ns/op	    9767 B/op	      31 allocs/op
BenchmarkCacheGetRandomKeyInt/redis-4                	     531	   2271602 ns/op	    9767 B/op	      31 allocs/op
BenchmarkCacheGetRandomKeyInt/redis-8                	     522	   2273678 ns/op	    9767 B/op	      31 allocs/op
BenchmarkCacheGetRandomKeyInt/redis-16               	     552	   2180911 ns/op	    9767 B/op	      31 allocs/op
BenchmarkCacheGetRandomKeyInt/pebble
BenchmarkCacheGetRandomKeyInt/pebble                 	  752023	      1575 ns/op	    1359 B/op	       7 allocs/op
BenchmarkCacheGetRandomKeyInt/pebble-2               	  699300	      1557 ns/op	    1359 B/op	       7 allocs/op
BenchmarkCacheGetRandomKeyInt/pebble-4               	  730688	      1534 ns/op	    1359 B/op	       7 allocs/op
BenchmarkCacheGetRandomKeyInt/pebble-8               	  768183	      1508 ns/op	    1359 B/op	       7 allocs/op
BenchmarkCacheGetRandomKeyInt/pebble-16              	  735848	      1506 ns/op	    1359 B/op	       7 allocs/op
BenchmarkCacheGetStruct
BenchmarkCacheGetStruct/encoding
BenchmarkCacheGetStruct/encoding                     	 2252955	       524 ns/op	     208 B/op	       4 allocs/op
BenchmarkCacheGetStruct/encoding-2                   	 2332430	       515 ns/op	     208 B/op	       4 allocs/op
BenchmarkCacheGetStruct/encoding-4                   	 2251696	       525 ns/op	     208 B/op	       4 allocs/op
BenchmarkCacheGetStruct/encoding-8                   	 2235301	       520 ns/op	     208 B/op	       4 allocs/op
BenchmarkCacheGetStruct/encoding-16                  	 2224682	       527 ns/op	     208 B/op	       4 allocs/op
BenchmarkCacheGetStruct/map
BenchmarkCacheGetStruct/map                          	 8009500	       141 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetStruct/map-2                        	 8406175	       143 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetStruct/map-4                        	 8249924	       145 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetStruct/map-8                        	 8324671	       145 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetStruct/map-16                       	 8102042	       145 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetStruct/shard
BenchmarkCacheGetStruct/shard                        	 7179788	       164 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetStruct/shard-2                      	 7332114	       164 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetStruct/shard-4                      	 6999268	       174 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetStruct/shard-8                      	 7028054	       170 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetStruct/shard-16                     	 6986014	       170 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetStruct/lru
BenchmarkCacheGetStruct/lru                          	 5818656	       207 ns/op	      48 B/op	       1 allocs/op
BenchmarkCacheGetStruct/lru-2                        	 5859214	       204 ns/op	      48 B/op	       1 allocs/op
BenchmarkCacheGetStruct/lru-4                        	 5518066	       210 ns/op	      48 B/op	       1 allocs/op
BenchmarkCacheGetStruct/lru-8                        	 5618907	       209 ns/op	      48 B/op	       1 allocs/op
BenchmarkCacheGetStruct/lru-16                       	 5617592	       214 ns/op	      48 B/op	       1 allocs/op
BenchmarkCacheGetStruct/ristretto
BenchmarkCacheGetStruct/ristretto                    	 7409641	       158 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetStruct/ristretto-2                  	 6809439	       175 ns/op	       7 B/op	       0 allocs/op
BenchmarkCacheGetStruct/ristretto-4                  	 6004058	       194 ns/op	       7 B/op	       0 allocs/op
BenchmarkCacheGetStruct/ristretto-8                  	 6170220	       192 ns/op	       7 B/op	       0 allocs/op
BenchmarkCacheGetStruct/ristretto-16                 	 6219170	       190 ns/op	       7 B/op	       0 allocs/op
BenchmarkCacheGetStruct/memcache
BenchmarkCacheGetStruct/memcache                     	    1412	    801366 ns/op	     424 B/op	      14 allocs/op
BenchmarkCacheGetStruct/memcache-2                   	    1395	    845730 ns/op	     424 B/op	      14 allocs/op
BenchmarkCacheGetStruct/memcache-4                   	    1454	    754811 ns/op	     424 B/op	      14 allocs/op
BenchmarkCacheGetStruct/memcache-8                   	    1509	    754192 ns/op	     425 B/op	      14 allocs/op
BenchmarkCacheGetStruct/memcache-16                  	    1354	    800273 ns/op	     428 B/op	      14 allocs/op
BenchmarkCacheGetStruct/redis
BenchmarkCacheGetStruct/redis                        	     553	   2131603 ns/op	    9720 B/op	      32 allocs/op
BenchmarkCacheGetStruct/redis-2                      	     548	   2139096 ns/op	    9720 B/op	      32 allocs/op
BenchmarkCacheGetStruct/redis-4                      	     537	   2211997 ns/op	    9720 B/op	      32 allocs/op
BenchmarkCacheGetStruct/redis-8                      	     524	   2189316 ns/op	    9720 B/op	      32 allocs/op
BenchmarkCacheGetStruct/redis-16                     	     548	   2185637 ns/op	    9720 B/op	      32 allocs/op
BenchmarkCacheGetStruct/pebble
BenchmarkCacheGetStruct/pebble                       	 1427671	       796 ns/op	    1248 B/op	       7 allocs/op
BenchmarkCacheGetStruct/pebble-2                     	 1448547	       830 ns/op	    1248 B/op	       7 allocs/op
BenchmarkCacheGetStruct/pebble-4                     	 1405844	       835 ns/op	    1248 B/op	       7 allocs/op
BenchmarkCacheGetStruct/pebble-8                     	 1441484	       831 ns/op	    1248 B/op	       7 allocs/op
BenchmarkCacheGetStruct/pebble-16                    	 1387006	       827 ns/op	    1248 B/op	       7 allocs/op
BenchmarkCacheSetStruct
BenchmarkCacheSetStruct/encoding
BenchmarkCacheSetStruct/encoding                     	 1000000	      1625 ns/op	     651 B/op	       5 allocs/op
BenchmarkCacheSetStruct/encoding-2                   	 1720123	       945 ns/op	     457 B/op	       5 allocs/op
BenchmarkCacheSetStruct/encoding-4                   	 1669809	       705 ns/op	     183 B/op	       4 allocs/op
BenchmarkCacheSetStruct/encoding-8                   	 1657442	       706 ns/op	     183 B/op	       4 allocs/op
BenchmarkCacheSetStruct/encoding-16                  	 1648228	       709 ns/op	     184 B/op	       4 allocs/op
BenchmarkCacheSetStruct/map
BenchmarkCacheSetStruct/map                          	 1000000	      1280 ns/op	     410 B/op	       9 allocs/op
BenchmarkCacheSetStruct/map-2                        	 1878842	      1517 ns/op	     341 B/op	       7 allocs/op
BenchmarkCacheSetStruct/map-4                        	 1790534	       692 ns/op	     263 B/op	       6 allocs/op
BenchmarkCacheSetStruct/map-8                        	 1792663	       665 ns/op	     263 B/op	       6 allocs/op
BenchmarkCacheSetStruct/map-16                       	 1762833	       677 ns/op	     263 B/op	       6 allocs/op
BenchmarkCacheSetStruct/shard
BenchmarkCacheSetStruct/shard                        	 1000000	      1437 ns/op	     411 B/op	       9 allocs/op
BenchmarkCacheSetStruct/shard-2                      	 1716608	       830 ns/op	     346 B/op	       7 allocs/op
BenchmarkCacheSetStruct/shard-4                      	 1647408	       736 ns/op	     263 B/op	       6 allocs/op
BenchmarkCacheSetStruct/shard-8                      	 1657657	       710 ns/op	     263 B/op	       6 allocs/op
BenchmarkCacheSetStruct/shard-16                     	 1651122	       711 ns/op	     263 B/op	       6 allocs/op
BenchmarkCacheSetStruct/lru
BenchmarkCacheSetStruct/lru                          	 1669929	       717 ns/op	     330 B/op	       8 allocs/op
BenchmarkCacheSetStruct/lru-2                        	 1666970	       686 ns/op	     330 B/op	       8 allocs/op
BenchmarkCacheSetStruct/lru-4                        	 1569268	       707 ns/op	     330 B/op	       8 allocs/op
BenchmarkCacheSetStruct/lru-8                        	 1569517	       701 ns/op	     330 B/op	       8 allocs/op
BenchmarkCacheSetStruct/lru-16                       	 1569993	       720 ns/op	     330 B/op	       8 allocs/op
BenchmarkCacheSetStruct/ristretto
BenchmarkCacheSetStruct/ristretto                    	 1665415	      1203 ns/op	     406 B/op	       5 allocs/op
BenchmarkCacheSetStruct/ristretto-2                  	 1000000	      1111 ns/op	     325 B/op	       5 allocs/op
BenchmarkCacheSetStruct/ristretto-4                  	 1000000	      1204 ns/op	     319 B/op	       5 allocs/op
BenchmarkCacheSetStruct/ristretto-8                  	 1000000	      1193 ns/op	     319 B/op	       5 allocs/op
BenchmarkCacheSetStruct/ristretto-16                 	  946750	      1171 ns/op	     324 B/op	       5 allocs/op
BenchmarkCacheSetStruct/memcache
BenchmarkCacheSetStruct/memcache                     	    1572	    733672 ns/op	     286 B/op	       9 allocs/op
BenchmarkCacheSetStruct/memcache-2                   	    1341	    799704 ns/op	     286 B/op	       9 allocs/op
BenchmarkCacheSetStruct/memcache-4                   	    1492	    810459 ns/op	     287 B/op	       9 allocs/op
BenchmarkCacheSetStruct/memcache-8                   	    1500	    807919 ns/op	     289 B/op	       9 allocs/op
BenchmarkCacheSetStruct/memcache-16                  	    1598	    773923 ns/op	     290 B/op	       9 allocs/op
BenchmarkCacheSetStruct/redis
BenchmarkCacheSetStruct/redis                        	     848	   1312946 ns/op	    9788 B/op	      35 allocs/op
BenchmarkCacheSetStruct/redis-2                      	     834	   1370112 ns/op	    9789 B/op	      35 allocs/op
BenchmarkCacheSetStruct/redis-4                      	     858	   1367748 ns/op	    9789 B/op	      35 allocs/op
BenchmarkCacheSetStruct/redis-8                      	     906	   1348890 ns/op	    9790 B/op	      35 allocs/op
BenchmarkCacheSetStruct/redis-16                     	     856	   1377737 ns/op	    9791 B/op	      35 allocs/op
BenchmarkCacheSetStruct/pebble
BenchmarkCacheSetStruct/pebble                       	     172	   6891869 ns/op	     179 B/op	       4 allocs/op
BenchmarkCacheSetStruct/pebble-2                     	     176	   7100201 ns/op	     189 B/op	       4 allocs/op
BenchmarkCacheSetStruct/pebble-4                     	     176	   6765299 ns/op	     417 B/op	       4 allocs/op
BenchmarkCacheSetStruct/pebble-8                     	     174	   6709812 ns/op	     196 B/op	       4 allocs/op
BenchmarkCacheSetStruct/pebble-16                    	     176	   6872531 ns/op	     207 B/op	       4 allocs/op
BenchmarkCacheGetParallel
BenchmarkCacheGetParallel/encoding
BenchmarkCacheGetParallel/encoding                   	 3755816	       393 ns/op	     192 B/op	       2 allocs/op
BenchmarkCacheGetParallel/encoding-2                 	 6620756	       200 ns/op	     192 B/op	       2 allocs/op
BenchmarkCacheGetParallel/encoding-4                 	10706964	       126 ns/op	     192 B/op	       2 allocs/op
BenchmarkCacheGetParallel/encoding-8                 	15889144	        83.4 ns/op	     192 B/op	       2 allocs/op
BenchmarkCacheGetParallel/encoding-16                	18838454	        67.2 ns/op	     192 B/op	       2 allocs/op
BenchmarkCacheGetParallel/map
BenchmarkCacheGetParallel/map                        	 8287477	       137 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetParallel/map-2                      	11197053	       101 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetParallel/map-4                      	19310756	        58.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetParallel/map-8                      	28979271	        37.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetParallel/map-16                     	40122621	        25.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetParallel/shard
BenchmarkCacheGetParallel/shard                      	 6388084	       175 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetParallel/shard-2                    	 9824578	       119 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetParallel/shard-4                    	16162353	        70.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetParallel/shard-8                    	23337940	        45.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetParallel/shard-16                   	34489749	        31.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetParallel/lru
BenchmarkCacheGetParallel/lru                        	 5497556	       219 ns/op	      48 B/op	       1 allocs/op
BenchmarkCacheGetParallel/lru-2                      	 5108966	       239 ns/op	      48 B/op	       1 allocs/op
BenchmarkCacheGetParallel/lru-4                      	 4236541	       277 ns/op	      48 B/op	       1 allocs/op
BenchmarkCacheGetParallel/lru-8                      	 3867518	       313 ns/op	      48 B/op	       1 allocs/op
BenchmarkCacheGetParallel/lru-16                     	 3719572	       323 ns/op	      48 B/op	       1 allocs/op
BenchmarkCacheGetParallel/ristretto
BenchmarkCacheGetParallel/ristretto                  	 6272048	       170 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetParallel/ristretto-2                	10652374	       103 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetParallel/ristretto-4                	15653863	        73.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetParallel/ristretto-8                	17346794	        64.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetParallel/ristretto-16               	18895278	        57.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkCacheGetParallel/memcache
BenchmarkCacheGetParallel/memcache                   	    1398	    868052 ns/op	     432 B/op	      12 allocs/op
BenchmarkCacheGetParallel/memcache-2                 	    2768	    446176 ns/op	     432 B/op	      12 allocs/op
BenchmarkCacheGetParallel/memcache-4                 	    4094	    244463 ns/op	     439 B/op	      12 allocs/op
BenchmarkCacheGetParallel/memcache-8                 	    8526	    141027 ns/op	     642 B/op	      12 allocs/op
BenchmarkCacheGetParallel/memcache-16                	   10852	    108739 ns/op	     871 B/op	      13 allocs/op
BenchmarkCacheGetParallel/redis
BenchmarkCacheGetParallel/redis                      	     524	   2255655 ns/op	    9704 B/op	      30 allocs/op
BenchmarkCacheGetParallel/redis-2                    	     933	   1244186 ns/op	    9704 B/op	      30 allocs/op
BenchmarkCacheGetParallel/redis-4                    	    1560	    790918 ns/op	    9704 B/op	      30 allocs/op
BenchmarkCacheGetParallel/redis-8                    	    1885	    613956 ns/op	    9704 B/op	      30 allocs/op
BenchmarkCacheGetParallel/redis-16                   	    2102	    558509 ns/op	    9704 B/op	      30 allocs/op
BenchmarkCacheGetParallel/pebble
BenchmarkCacheGetParallel/pebble                     	 1377561	       970 ns/op	    1248 B/op	       5 allocs/op
BenchmarkCacheGetParallel/pebble-2                   	 2307592	       630 ns/op	    1248 B/op	       5 allocs/op
BenchmarkCacheGetParallel/pebble-4                   	 3651379	       352 ns/op	    1248 B/op	       5 allocs/op
BenchmarkCacheGetParallel/pebble-8                   	 5771799	       222 ns/op	    1248 B/op	       5 allocs/op
BenchmarkCacheGetParallel/pebble-16                  	 7370930	       187 ns/op	    1248 B/op	       5 allocs/op
BenchmarkCacheSetParallel
BenchmarkCacheSetParallel/encoding
BenchmarkCacheSetParallel/encoding                   	 2877936	       417 ns/op	     176 B/op	       5 allocs/op
BenchmarkCacheSetParallel/encoding-2                 	 3406563	       377 ns/op	     176 B/op	       5 allocs/op
BenchmarkCacheSetParallel/encoding-4                 	 3595508	       334 ns/op	     176 B/op	       5 allocs/op
BenchmarkCacheSetParallel/encoding-8                 	 3169240	       378 ns/op	     176 B/op	       5 allocs/op
BenchmarkCacheSetParallel/encoding-16                	 3012076	       400 ns/op	     176 B/op	       5 allocs/op
BenchmarkCacheSetParallel/map
BenchmarkCacheSetParallel/map                        	 2573331	       523 ns/op	     256 B/op	       7 allocs/op
BenchmarkCacheSetParallel/map-2                      	 3006007	       453 ns/op	     256 B/op	       7 allocs/op
BenchmarkCacheSetParallel/map-4                      	 3168489	       392 ns/op	     256 B/op	       7 allocs/op
BenchmarkCacheSetParallel/map-8                      	 2839058	       440 ns/op	     256 B/op	       7 allocs/op
BenchmarkCacheSetParallel/map-16                     	 2834168	       431 ns/op	     256 B/op	       7 allocs/op
BenchmarkCacheSetParallel/shard
BenchmarkCacheSetParallel/shard                      	 2250218	       516 ns/op	     256 B/op	       7 allocs/op
BenchmarkCacheSetParallel/shard-2                    	 3581533	       370 ns/op	     256 B/op	       7 allocs/op
BenchmarkCacheSetParallel/shard-4                    	 3066703	       415 ns/op	     256 B/op	       7 allocs/op
BenchmarkCacheSetParallel/shard-8                    	 2774422	       428 ns/op	     256 B/op	       7 allocs/op
BenchmarkCacheSetParallel/shard-16                   	 2749574	       432 ns/op	     256 B/op	       7 allocs/op
BenchmarkCacheSetParallel/lru
BenchmarkCacheSetParallel/lru                        	 3272673	       430 ns/op	     240 B/op	       6 allocs/op
BenchmarkCacheSetParallel/lru-2                      	 4692276	       278 ns/op	     240 B/op	       6 allocs/op
BenchmarkCacheSetParallel/lru-4                      	 3994620	       312 ns/op	     240 B/op	       6 allocs/op
BenchmarkCacheSetParallel/lru-8                      	 3531354	       341 ns/op	     240 B/op	       6 allocs/op
BenchmarkCacheSetParallel/lru-16                     	 3414451	       353 ns/op	     240 B/op	       6 allocs/op
BenchmarkCacheSetParallel/ristretto
BenchmarkCacheSetParallel/ristretto                  	 2669528	       456 ns/op	     224 B/op	       5 allocs/op
BenchmarkCacheSetParallel/ristretto-2                	 2214732	       547 ns/op	     224 B/op	       5 allocs/op
BenchmarkCacheSetParallel/ristretto-4                	 2122172	       564 ns/op	     224 B/op	       5 allocs/op
BenchmarkCacheSetParallel/ristretto-8                	 1858959	       639 ns/op	     224 B/op	       5 allocs/op
BenchmarkCacheSetParallel/ristretto-16               	 1821427	       656 ns/op	     224 B/op	       5 allocs/op
BenchmarkCacheSetParallel/memcache
BenchmarkCacheSetParallel/memcache                   	 1395469	       863 ns/op	     352 B/op	       8 allocs/op
BenchmarkCacheSetParallel/memcache-2                 	 1672177	       705 ns/op	     352 B/op	       8 allocs/op
BenchmarkCacheSetParallel/memcache-4                 	  848569	      1414 ns/op	     406 B/op	       8 allocs/op
BenchmarkCacheSetParallel/memcache-8                 	  742070	      1361 ns/op	     402 B/op	       8 allocs/op
BenchmarkCacheSetParallel/memcache-16                	 1346508	       950 ns/op	     353 B/op	       8 allocs/op
BenchmarkCacheSetParallel/redis
BenchmarkCacheSetParallel/redis                      	     100	  21280373 ns/op	    1526 B/op	      31 allocs/op
BenchmarkCacheSetParallel/redis-2                    	     908	   1203995 ns/op	    1520 B/op	      31 allocs/op
BenchmarkCacheSetParallel/redis-4                    	     901	   1171409 ns/op	    1520 B/op	      31 allocs/op
BenchmarkCacheSetParallel/redis-8                    	     948	   1185400 ns/op	    1521 B/op	      31 allocs/op
BenchmarkCacheSetParallel/redis-16                   	     852	   1247485 ns/op	    1523 B/op	      31 allocs/op
BenchmarkCacheSetParallel/pebble
BenchmarkCacheSetParallel/pebble                     	     174	   6865631 ns/op	     178 B/op	       5 allocs/op
BenchmarkCacheSetParallel/pebble-2                   	     189	   6658668 ns/op	     397 B/op	       5 allocs/op
BenchmarkCacheSetParallel/pebble-4                   	     360	   3477886 ns/op	     184 B/op	       5 allocs/op
BenchmarkCacheSetParallel/pebble-8                   	     717	   1716858 ns/op	     184 B/op	       5 allocs/op
BenchmarkCacheSetParallel/pebble-16                  	    1417	    956456 ns/op	     182 B/op	       5 allocs/op
```
