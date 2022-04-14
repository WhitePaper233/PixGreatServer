# PixGreatServer
The backend API server of PixGreat repo

## Performance:
`Test environment:`
```bash
CPU: AMD Ryzen 5 2600X - 3.60 GHz
RAM: DDR4-3000 8G * 2 (16G Dual-Channel)
SYSTEM: Windows 10 Professional Workstation 21H2 - 19044
COMPILER: Go 1.18
BUILD COMMAND: go build
TEST COMMAND: wrk -t100 -c1000 -d30s  --latency "http://127.0.0.1:8080"
```
`With MetadataLoadToMem optimize OFF:`
```
Running 30s test @ http://127.0.0.1:8080
  100 threads and 1000 connections
  Thread Stats   Avg      Stdev    Max     +/- Stdev
    Latency    73.60ms   43.58ms 563.08ms   71.45%
    Req/Sec    140.36    40.44   313.00     70.22%
  Latency Distribution
     50%   70.11ms
     75%   77.17ms
     90%  136.50ms
     99%  212.63ms
  421019 requests in 30.10s, 121.02MB read
  Socket errors: connect 0, read 0, write 1066, timeout 0
Requests/sec:  13987.74
Transfer/sec:  4.02MB
```

`With MetadataLoadToMem optimize ON:`
```
Running 30s test @ http://127.0.0.1:8080
  100 threads and 1000 connections
  Thread Stats   Avg      Stdev    Max     +/- Stdev
    Latency    67.15ms   41.62ms 462.17ms   69.28%
    Req/Sec    154.42    42.23   330.00     62.90%
  Latency Distribution
     50%   64.40ms
     75%   69.04ms
     90%  126.91ms
     99%  197.62ms
  462774 requests in 30.10s, 133.01MB read
  Socket errors: connect 0, read 0, write 1250, timeout 0
Requests/sec:  15374.40
Transfer/sec:  4.42MB
```