#!/bin/bash
set -e

wd=$(realpath $(dirname $0)/..)
export BUILD_DIR=$(realpath $(dirname $0)/..)
pushd $BUILD_DIR > /dev/null
export GOVERSION=${GOVERSION:-1.11}
export REVISION=${GIT_COMMIT:-$(git rev-parse HEAD)}
export LIBRARY_REVISION="3f4dacd4d38a89e715d8c5f52f7cfbf776728a2f"
export VERSION=$(cat ${wd}/.version)
export SOURCE_DATE_EPOCH=$(git log -1 --format=%ct)
export LAST_CHANGE=$(git log -1 --format=%cd)

if [[ -z $LIBRARY_REVISION ]]; then
  echo "unable to get library revision"
  exit 1
fi
