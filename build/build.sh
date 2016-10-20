#!/bin/bash
set -e

wd=$(realpath $(dirname $0)/..)
source ${wd}/build/config.sh

export DIR=$(mktemp -d)

tar --create . | docker run --rm -i golang:$GOVERSION bash -c "$(cat build/docker_build.sh)" > $DIR/build.tar

cd $DIR

tar --extract --file=build.tar
rm -f build.tar

# Homebrew - binary must be called phraseapp, because the binary name inside
# the tar will be made available system wide
cp phraseapp_macosx_amd64 phraseapp
tar cfz phraseapp_macosx_amd64.tar.gz phraseapp
rm phraseapp

for name in phraseapp_linux_386 phraseapp_linux_amd64; do
  tar cfz ${name}.tar.gz $name
done

if ! which zip > /dev/null; then
  apt-get update && apt-get install -y zip
fi

zip phraseapp_windows_amd64.exe.zip phraseapp_windows_amd64.exe > /dev/null

echo $DIR > ${wd}/.bin_dir
echo -n $BUILD_VERSION > ${DIR}/.build_version
