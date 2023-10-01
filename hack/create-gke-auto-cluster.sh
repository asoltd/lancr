#!/bin/bash

gcloud container clusters create-auto lancr-prod \
  --project=asoltd-guilds \
  --region=europe-central2
