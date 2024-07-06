#!/usr/bin/bash

cd "$(dirname "$0")" || exit 1

cd ..

# shellcheck disable=SC2164
cd app/user

sudo docker build -t mutezebra/tiktok-user:2.0 .

# shellcheck disable=SC2164
cd ../../scripts/images

sudo docker save -o user.tar mutezebra/tiktok-user:2.0

sudo ctr -n=k8s.io i import user.tar
sleep 3
cd ../..

sudo kubectl delete -f deploy/user/user.yaml
sudo kubectl apply -f deploy/user/config.yaml
sleep 3
sudo kubectl apply -f deploy/user/user.yaml
