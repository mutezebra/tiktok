#!/usr/bin/bash

cd "$(dirname "$0")" || exit 1

# 回到项目根目录
cd ..

# shellcheck disable=SC2164
cd app/gateway

# 将 gateway 模块构建为镜像
sudo docker build -t mutezebra/tiktok-gateway:2.0 .

# shellcheck disable=SC2164
cd ../../scripts/images

# 将刚刚构建好的镜像保存为 tar 格式
sudo docker save -o gateway.tar mutezebra/tiktok-gateway:2.0

# 将 tar 格式的镜像文件导入到 k8s 中
sudo ctr -n=k8s.io i import gateway.tar

# 回到项目根目录
cd ../..

# shellcheck disable=SC2164
sudo kubectl delete -f deploy/gateway/gateway.yaml
sudo kubectl apply -f deploy/gateway/config.yaml

# 等待相关配置的部署
sleep 2
sudo kubectl apply -f deploy/gateway/gateway.yaml
