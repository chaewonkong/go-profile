#!/bin/bash

echo "Load Test"


cleanup() {
    echo "Cleaning up..."
    # Kill all child processes in the same process group
    kill 0
}

# Trap SIGINT (Ctrl + C) to run the cleanup function
trap cleanup SIGINT


# 프로파일링
curl http://localhost:8080/debug/pprof/trace\?seconds\=30 --output trace.out &
curl http://localhost:8080/debug/pprof/heap\?seconds\=30 --output heap.out &
curl http://localhost:8080/debug/pprof/profile\? seconds\=30 --output cpu.out &
# 부하를 준 상태에서 GET 요청 수행 

# recursive
bombardier -t 1s -l -c 10 -d 30s \
-H "Content-Type: application/json" \
-b "$PAYLOAD" \
-m GET http://localhost:8080/fib/recursive/30 &

# using cache
bombardier -t 1s -l -c 10 -d 30s \
-H "Content-Type: application/json" \
-b "$PAYLOAD" \
-m GET http://localhost:8080/fib/cache/30000 &

wait
# go tool 활성화 -> port에서 각 프로파일 실행
go tool trace -http localhost:9091 trace.out &
go tool pprof -http localhost:9092 heap.out &
go tool pprof -http localhost:9093 cpu.out &

wait