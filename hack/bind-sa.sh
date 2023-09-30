#!/bin/bash

GCP_SERVICE_ACCOUNT=firestore-editor@asoltd-guilds.iam.gserviceaccount.com

kubectl annotate serviceaccount heroes-sa \
--namespace default \
iam.gke.io/gcp-service-account=$GCP_SERVICE_ACCOUNT

gcloud iam service-accounts add-iam-policy-binding \
$GCP_SERVICE_ACCOUNT \
--role=roles/iam.workloadIdentityUser \
--member="serviceAccount:asoltd-guilds.svc.id.goog[default/heroes-sa]" \
--project=asoltd-guilds