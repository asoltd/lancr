#!/bin/bash

gcloud container clusters update lancr-prod \
--region=europe-central2 \
--workload-pool=asoltd-guilds.svc.id.goog \
--project=asoltd-guilds
