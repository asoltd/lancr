#!/bin/bash

gcloud container clusters update lancr-cluster \
--region=europe-central2 \
--workload-pool=asoltd-guilds.svc.id.goog \
--project=asoltd-guilds
