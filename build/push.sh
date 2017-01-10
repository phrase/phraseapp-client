#!/bin/bash
set -e

WD=$(realpath $(dirname $0)/..)
BIN_DIR=$(cat ${WD}/.bin_dir)
BUILD_VERSION=$(cat $BIN_DIR/.build_version)

if ! which aws > /dev/null; then
  if which yum > /dev/null 2>&1 ; then
    yum install -y python2-pip
  else
    apt-get update
    apt-get install -y python-pip
  fi
  pip install awscli
fi

# probably running inside jenkins
dst=s3://phraseapp-client-releases/${BUILD_VERSION}
for p in $(find $BIN_DIR -type f); do
  shasum=$(sha256sum $p | awk '{ print $1 }')
  if [[ -z $shasum ]]; then
    echo "unable to get shasum of $p"
    exit 1
  fi
  aws s3 cp --acl=public-read --metadata SHA256=${shasum} $p  ${dst}/$(basename $p)
done
