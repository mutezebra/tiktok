#!/usr/bin/bash

# 相关操作注释请见 gateway.sh
cd "$(dirname "$0")" || exit 1

cd ..

# 将 video 模块构建为镜像
sudo docker build -t mutezebra/tiktok-video:2.0 --build-arg SERVICE=video --build-arg PORT=10002 .

# 检查容器运行时是否为 docker
runtime=$(kubectl get nodes -o jsonpath='{.items[0].status.nodeInfo.containerRuntimeVersion}')

# 判断运行时是否是 docker
if [[ "$runtime" != docker://* ]]; then
  cd scripts/images || exit

  # 将刚刚构建好的镜像保存为 tar 格式
  sudo docker save -o video.tar mutezebra/tiktok-video:2.0

  # 将 tar 格式的镜像文件导入到 k8s 中
  sudo ctr -n=k8s.io i import video.tar

  sudo rm -f video.tar

  # 回到项目根目录
  cd ../..
fi

sudo kubectl delete -f deploy/video/video.yaml
sudo kubectl apply -f deploy/video/config.yaml
sleep 3
sudo kubectl apply -f deploy/video/video.yaml
