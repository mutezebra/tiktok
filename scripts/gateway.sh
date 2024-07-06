#!/usr/bin/bash

cd "$(dirname "$0")" || exit 1

cd ..

# shellcheck disable=SC2164
cd app/gateway

sudo docker build -t mutezebra/tiktok-gateway:2.0 .

# shellcheck disable=SC2164
cd ../../scripts/images

sudo docker save -o gateway.tar mutezebra/tiktok-gateway:2.0

sudo ctr -n=k8s.io i import gateway.tar

cd ../..

# shellcheck disable=SC2164
sudo kubectl delete -f deploy/gateway/gateway.yaml
sudo kubectl apply -f deploy/gateway/config.yaml
sleep 2
sudo kubectl apply -f deploy/gateway/gateway.yaml
