#!/bin/bash

gcloud deploy apply \
--file=clouddeploy.yaml \
--region=europe-central2 \
--project=asoltd-guilds