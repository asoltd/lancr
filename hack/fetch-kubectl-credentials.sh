#!/bin/bash

gcloud container clusters get-credentials lancr-prod \
  --region=europe-central2 \
  --project=asoltd-guilds
