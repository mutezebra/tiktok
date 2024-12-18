#!/usr/bin/bash

# 定义检查 Pod 状态的函数
function wait_for_pods_running() {
  local LABEL_SELECTOR=$1
  local CHECK_INTERVAL=${2:-5}  # 默认每 5 秒检查一次
  local MAX_WAIT_TIME=${3:-300} # 默认最大等待时间为 300 秒
  local MODULE_NAME=${4:-"unnamed module"}
  local WAIT_TIME=0
  local COLOR_END="\033[0m"
  local GREEN_COLOR_START="\033[32m"
  local RED_COLOR_START="\033[31m"

  echo "$GREEN_COLOR_START Waiting for $MODULE_NAME pods to be in Running state... $COLOR_END"

  while true; do
    # 获取 Pod 的状态
    POD_STATUSES=$(kubectl get pods -l "$LABEL_SELECTOR" -o jsonpath='{.items[*].status.phase}')

    # 检查是否所有 Pod 都处于 Running 状态
    RUNNING_COUNT=$(echo "$POD_STATUSES" | grep -w "Running" -c)
    TOTAL_COUNT=$(echo "$POD_STATUSES" | wc -w)

    if [ "$RUNNING_COUNT" -eq "$TOTAL_COUNT" ]; then
      echo "All Pods are in Running state."
      break
    fi

    # 检查是否超过最大等待时间
    if [ $WAIT_TIME -ge "$MAX_WAIT_TIME" ]; then
      echo "$RED_COLOR_START Timeout: Pods did not reach Running state within $MAX_WAIT_TIME seconds. $COLOR_END"
      exit 1
    fi

    # 等待 CHECK_INTERVAL 秒后再次检查
    echo "Pods are not ready yet. Retrying in $CHECK_INTERVAL seconds..."
    sleep "$CHECK_INTERVAL"
    WAIT_TIME=$((WAIT_TIME + CHECK_INTERVAL))
  done
}

cd "$(dirname "$0")" || exit 1

# 回到项目根目录
cd ..

# 创建 全局容器的 pv
sudo kubectl apply -f deploy/common/etcd-pv.yaml
sudo kubectl apply -f deploy/common/initdb.yaml
sleep 2

# 创建容器, 如果没有的话需要等待
sudo kubectl apply -f deploy/common/etcd.yaml
sudo kubectl apply -f deploy/common/jaeger.yaml
sudo kubectl apply -f deploy/common/mysql.yaml

wait_for_pods_running "app=etcd-pod,app=jaeger-pod,app=mysql-pod" 5 300 "global"

# 创建 gateway 模块的 pv
sudo kubectl apply -f deploy/gateway/kafka-pv.yaml
sleep 2

sudo kubectl apply -f deploy/gateway/kafka.yaml
sudo kubectl apply -f deploy/gateway/kafka-ui.yaml

# 等待
wait_for_pods_running "app=gateway-kafka-pod,app=kafka-ui" 5 150 "gateway"

# 创建 video 模块的 pv
sudo kubectl apply -f deploy/video/redis-pv.yaml
sleep 2

sudo kubectl apply -f deploy/video/redis.yaml

# 等待
wait_for_pods_running "app=tiktok-video-redis-pod" 5 150 "video"
