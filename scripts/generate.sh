#!/bin/bash

set -euo pipefail

docker build -t vss:latest .
docker create --name dummy vss:latest
docker cp dummy:/out/vss_schema_docs .
docker cp dummy:/out/data_schema_struct .
docker rm -f dummy
