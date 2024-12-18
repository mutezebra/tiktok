#!/usr/bin/bash

# 相关操作注释请见 gateway.sh
cd "$(dirname "$0")" || exit 1

cd ..

# 将 relation 模块构建为镜像
sudo docker build -t mutezebra/tiktok-relation:2.0 --build-arg SERVICE=relation --build-arg PORT=10004 .

# 检查容器运行时是否为 docker
runtime=$(kubectl get nodes -o jsonpath='{.items[0].status.nodeInfo.containerRuntimeVersion}')

# 判断运行时是否是 docker
if [[ "$runtime" != docker://* ]]; then
  cd scripts/images || exit

  # 将刚刚构建好的镜像保存为 tar 格式
  sudo docker save -o relation.tar mutezebra/tiktok-relation:2.0

  # 将 tar 格式的镜像文件导入到 k8s 中
  sudo ctr -n=k8s.io i import relation.tar

  sudo rm -f relation.tar

  # 回到项目根目录
  cd ../..
fi

sudo kubectl delete -f deploy/relation/relation.yaml
sudo kubectl apply -f deploy/relation/config.yaml
sleep 3
sudo kubectl apply -f deploy/relation/relation.yaml
