#!/bin/bash

gcloud deploy releases create release-$(git rev-parse --short HEAD) \
  --delivery-pipeline=lancr-deployment \
  --region=europe-central2 \
  --source=./ \
  --images=lancr=gcr.io/asoltd-guilds/lancr:latest \
  --project=asoltd-guilds

