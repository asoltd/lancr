#!/bin/bash

ngrok tunnel \
  --label edge=edghts_2WzkgTRgLIuZR0lGpW89qr4G8GD \
  --region eu \
  http://$(minikube ip)

