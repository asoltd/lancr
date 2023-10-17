#!/bin/bash

docker build -t gcr.io/asoltd-guilds/lancr:latest . \
        && docker push gcr.io/asoltd-guilds/lancr:latest \
	&& skaffold apply k8s/manifest.yaml

