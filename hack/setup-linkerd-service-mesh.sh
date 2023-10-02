#!/bin/bash

gcloud container clusters update lancr-prod \
--workload-policies=allow-net-admin \
--project=asoltd-guilds \
--region=europe-central2

linkerd check --pre
