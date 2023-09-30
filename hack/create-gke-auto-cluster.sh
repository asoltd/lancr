#!/bin/bash

# cluster currently running name is lancr-cluster, there will be multiple
# environments, hence migrating the name to lancr-prod
gcloud container clusters create-auto \
lancr-prod \
--project=asoltd-guilds \
--region=europe-central2