#!/bin/bash

gcloud iam workload-identity-pools create gh-actions-pool \
--location=global \
--description="Pool for GitHub Actions" \
--project=asoltd-guilds

gcloud iam workload-identity-pools providers create-oidc gh-actions-provider \
--project=asoltd-guilds \
--location=global \
--workload-identity-pool=gh-actions-pool \
--display-name="Github Actions Provider" \
--attribute-mapping="google.subject=assertion.sub,attribute.actor=assertion.actor,attribute.repository=assertion.repository" \
--issuer-uri="https://token.actions.githubusercontent.com"