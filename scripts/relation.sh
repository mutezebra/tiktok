#!/usr/bin/bash

cd "$(dirname "$0")" || exit 1

cd ..

# shellcheck disable=SC2164
cd app/relation

sudo docker build -t mutezebra/tiktok-relation:2.0 .

# shellcheck disable=SC2164
cd ../../scripts/images

sudo docker save -o relation.tar mutezebra/tiktok-relation:2.0

sudo ctr -n=k8s.io i import relation.tar

cd ../..

sudo kubectl delete -f deploy/relation/relation.yaml
sudo kubectl apply -f deploy/relation/config.yaml
sleep 3
sudo kubectl apply -f deploy/relation/relation.yaml
