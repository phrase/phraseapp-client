#!/bin/bash
set -e

wd=$(realpath $(dirname $0)/..)
export BUILD_DIR=$(realpath $(dirname $0)/..)
pushd $BUILD_DIR > /dev/null
export GOVERSION=${GOVERSION:-1.10.3}
export REVISION=${GIT_COMMIT:-$(git rev-parse HEAD)}
export LIBRARY_REVISION=$(cat Gopkg.lock | grep github.com/phrase/phraseapp-go -A 2 | tail -n 1 | cut -d '"' -f 2)
export CURRENT_DATE=$(TZ=UTC date +"%Y-%m-%dT%H:%M:%SZ")
export VERSION=$(cat ${wd}/.version)

if [[ -z $LIBRARY_REVISION ]]; then
  echo "unable to get library revision"
  exit 1
fi

if [[ -z $VERSION ]]; then
  # try to fetch the most recent version and use <version>-dev
  export VERSION=$(git log --pretty=format:'%d' | grep -o 'tag: .*' | cut -d ')' -f 1 | awk '{ print $2 }' | sed -s 's/,$//g' | head -n 1)-dev
fi
