#!/bin/bash

gcloud compute addresses create lancr-lb-ip --global

gcloud compute addresses describe lancr-lb-ip --global

# address: 34.111.37.106