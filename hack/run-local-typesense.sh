#!/bin/bash

export TYPESENSE_API_KEY=xyz

docker run -p 8108:8108 \
            -v $HOME/typesense_data:/data typesense/typesense:0.25.1 \
            --data-dir /data \
            --api-key=$TYPESENSE_API_KEY \
            --enable-cors
#
