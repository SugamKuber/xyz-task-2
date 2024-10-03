#!/bin/sh

CQLSH_PATH="/app/venv/bin/cqlsh"

until $CQLSH_PATH -e "SHOW VERSION" 172.23.0.2 9042; do
  echo "Scylla is unavailable - sleeping"
  sleep 1
done

exec "/app/main"