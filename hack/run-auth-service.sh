#!/bin/bash

# if wanna build also 
# docker build -t gcr.io/asoltd/lancr:latest . 

LOCAL_CREDS_PATH=$HOME/.config/gcloud/application_default_credentials.json
CONTAINER_CREDS_PATH=/app/application_default_credentials.json

docker run -it --rm \
	-p 42069:42069 \
	-v $LOCAL_CREDS_PATH:$CONTAINER_CREDS_PATH \
	-e GOOGLE_APPLICATION_CREDENTIALS=$CONTAINER_CREDS_PATH \
	-d  \
	gcr.io/asoltd/lancr:latest \
	./lancr auth 

