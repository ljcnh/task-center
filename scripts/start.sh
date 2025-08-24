#!/bin/bash

# 启动API服务
cd cmd/api
go run main.go &
API_PID=$!

echo "API服务已启动，PID: $API_PID"

# 启动Worker服务
cd ../../cmd/worker
go run main.go &
WORKER_PID=$!

echo "Worker服务已启动，PID: $WORKER_PID"

# 等待退出信号
echo "按Ctrl+C停止所有服务..."

trap "kill $API_PID $WORKER_PID" SIGINT SIGTERM

wait

echo "所有服务已停止"