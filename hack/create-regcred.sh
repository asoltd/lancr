#!/bin/bash

kubectl create secret generic regcred \
    --from-file=.dockerconfigjson=~/.docker/config.json> \
    --type=kubernetes.io/dockerconfigjson
