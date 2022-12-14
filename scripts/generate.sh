#!/bin/bash

set -euo pipefail

docker build -t vss:latest .
docker create --name dummy vss:latest
docker cp dummy:/out .
docker rm -f dummy
