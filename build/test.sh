#!/bin/bash
set -e

source $(realpath $(dirname $0))/config.sh

tar --create . | docker run --rm -i golang:$GOVERSION bash -c "$(cat build/docker_test.sh)"
