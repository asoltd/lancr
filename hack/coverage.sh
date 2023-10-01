#!/bin/bash

set -e 

go test -v -coverprofile=coverage.out ./server
