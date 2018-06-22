#!/bin/bash

set -e

docker build -t phrase/innosetup innosetup/
docker run --rm -it --entrypoint /bin/bash -v $(pwd):/client phrase/innosetup /client/innosetup/docker_build.sh
