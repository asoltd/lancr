#!/bin/bash

PROJECT_ID=asoltd-guilds
SERVICE_ACCOUNT_NAME=auth-service-sa

# generate sa
gcloud iam service-accounts create $SERVICE_ACCOUNT_NAME \
	--display-name "Service Account for auth-service" \
	--project $PROJECT_ID

# generate key
gcloud iam service-accounts keys create $SERVICE_ACCOUNT_NAME.json \
	--iam-account $SERVICE_ACCOUNT_NAME@$PROJECT_ID.iam.gserviceaccount.com \
	--project $PROJECT_ID

# TODO way too wide permissions, need to narrow down
gcloud projects add-iam-policy-binding $PROJECT_ID \
	--member serviceAccount:$SERVICE_ACCOUNT_NAME@$PROJECT_ID.iam.gserviceaccount.com \
	--role roles/firebase.admin

kubectl create secret generic auth-service-sa \
	--from-file=$SERVICE_ACCOUNT_NAME.json \
	--namespace=default
