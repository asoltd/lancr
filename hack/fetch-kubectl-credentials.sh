#!/bin/bash

gcloud container clusters get-credentials lancr-cluster \
  --region=europe-central2 \
  --project=asoltd-guilds
