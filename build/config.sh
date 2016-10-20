#!/bin/bash
set -e

export BUILD_DIR=$(realpath $(dirname $0)/..)
pushd $BUILD_DIR > /dev/null
export GOVERSION=${GOVERSION:-1.7.1}
export REVISION=${GIT_COMMIT:-$(git rev-parse HEAD)}
export LIBRARY_REVISION=$(cat Godeps/Godeps.json | grep github.com/phrase/phraseapp-go -A 1 | tail -n 1 | cut -d '"' -f 4)
export PROJ_DIR=/go/src/github.com/phrase/phraseapp-client
export BUILD_VERSION=${VERSION:-$REVISION}

if [[ -z $LIBRARY_REVISION ]]; then
  echo "unable to get library revision"
  exit 1
fi

if [[ -z $VERSION ]]; then
  # try to fetch the most recent version and use <version>-dev
  export VERSION=$(git log --pretty=format:'%d' | grep -o 'tag: .*' | cut -d ')' -f 1 | awk '{ print $2 }' | sed -s 's/,$//g' | head -n 1)-dev
fi
