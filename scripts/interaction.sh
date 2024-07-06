#!/usr/bin/bash

cd "$(dirname "$0")" || exit 1

cd ..

# shellcheck disable=SC2164
cd app/interaction

sudo docker build -t mutezebra/tiktok-interaction:2.0 .

# shellcheck disable=SC2164
cd ../../scripts/images

sudo docker save -o interaction.tar mutezebra/tiktok-interaction:2.0

sudo ctr -n=k8s.io i import interaction.tar

cd ../..

sudo kubectl delete -f deploy/interaction/interaction.yaml
sudo kubectl apply -f deploy/interaction/config.yaml
sleep 3
sudo kubectl apply -f deploy/interaction/interaction.yaml
