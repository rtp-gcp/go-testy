#!/usr/bin/env bash

GITROOT=$(git rev-parse --show-toplevel)

export PATH=${GITROOT}/bin/:${PATH}

export GOOGLE_APPLICATION_CREDENTIALS="${GITROOT}/keys/shawdavis-244405e0b7c4.json"
