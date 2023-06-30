#!/usr/bin/env bash

# This script starts a local devnet using Docker Compose only
# Build by `make devnet-up-deploy`

. /etc/profile
set -eu

L1_URL="http://localhost:8545"
L2_URL="http://localhost:9545"

# Helper method that waits for a given URL to be up. Can't use
# cURL's built-in retry logic because connection reset errors
# are ignored unless you're using a very recent version of cURL
function wait_up {
  echo -n "Waiting for $1 to come up..."
  i=0
  until curl -s -f -o /dev/null "$1"
  do
    echo -n .
    sleep 0.25

    ((i=i+1))
    if [ "$i" -eq 300 ]; then
      echo " Timeout!" >&2
      exit 1
    fi
  done
  echo "Done!"
}

# rm -rf /tmp/l1_data /tmp/l2_data /tmp/op_log /tmp/postgres-data /tmp/postgres-data2 /tmp/redis-data /tmp/redis-data2


# Bring up L1.
(
  cd ops-bedrock
  echo "Bringing up L1..."
  docker-compose -f docker-compose-metis.yml up -d l1
  wait_up $L1_URL
)

# Bring up L2.
(
  cd ops-bedrock
  echo "Bringing up L2..."
  docker-compose -f docker-compose-metis.yml up -d l2
  wait_up $L2_URL
)

# dev, not deployed
L2OO_ADDRESS="0x6900000000000000000000000000000000000000"

# Bring up everything else.
(
  cd ops-bedrock
  echo "Bringing up devnet..."
  L2OO_ADDRESS="$L2OO_ADDRESS" \
      docker-compose -f docker-compose-metis.yml up -d op-node op-proposer op-batcher

  # echo "Bringing up stateviz webserver..."
  # docker-compose -f docker-compose-metis.yml up -d stateviz
)

echo "Devnet ready."
