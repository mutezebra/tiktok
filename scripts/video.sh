#!/usr/bin/bash

# 相关操作注释请见 gateway.sh
cd "$(dirname "$0")" || exit 1

cd ..

# shellcheck disable=SC2164
cd app/video

sudo docker build -t mutezebra/tiktok-video:2.0 .

# shellcheck disable=SC2164
cd ../../scripts/images

sudo docker save -o video.tar mutezebra/tiktok-video:2.0

sudo ctr -n=k8s.io i import video.tar

cd ../..

sudo kubectl delete -f deploy/video/video.yaml
sudo kubectl apply -f deploy/video/config.yaml
sleep 3
sudo kubectl apply -f deploy/video/video.yaml
