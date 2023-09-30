#!/bin/bash

gcloud deploy releases create release-$(git rev-parse --short HEAD) \
--delivery-pipeline=lancr-deployment \
--region=europe-central2 \
--source=./ \
--project=asoltd-guilds

