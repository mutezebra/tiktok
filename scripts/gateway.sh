#!/usr/bin/bash

cd "$(dirname "$0")" || exit 1

# 回到项目根目录
cd ..

# 将 gateway 模块构建为镜像
sudo docker build -t mutezebra/tiktok-gateway:2.0 --build-arg SERVICE=gateway --build-arg PORT=4000 .

# 检查容器运行时是否为 docker
runtime=$(kubectl get nodes -o jsonpath='{.items[0].status.nodeInfo.containerRuntimeVersion}')

# 判断运行时是否是 docker
if [[ "$runtime" != docker://* ]]; then
  cd scripts/images || exit

  # 将刚刚构建好的镜像保存为 tar 格式
  sudo docker save -o gateway.tar mutezebra/tiktok-gateway:2.0

  # 将 tar 格式的镜像文件导入到 k8s 中
  sudo ctr -n=k8s.io i import gateway.tar

  sudo rm -f gateway.tar

  # 回到项目根目录
  cd ../..
fi

# shellcheck disable=SC2164
sudo kubectl delete -f deploy/gateway/gateway.yaml
sudo kubectl apply -f deploy/gateway/config.yaml

# 等待相关配置的部署
sleep 2
sudo kubectl apply -f deploy/gateway/gateway.yaml
