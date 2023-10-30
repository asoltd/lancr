#!/bin/bash

# the command runs the process in the background
ngrok tunnel \
  --label edge=edghts_2WzkgTRgLIuZR0lGpW89qr4G8GD \
  --region eu \
  http://$(minikube ip) > /dev/null &

