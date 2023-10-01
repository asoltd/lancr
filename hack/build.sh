#!/bin/bash

go mod download

go build -ldflags="-s -w" .