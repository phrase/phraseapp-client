#!/bin/bash
set -e

wd=$(realpath $(dirname $0)/..)
source ${wd}/build/config.sh

export DIST_DIR=dist
rm -rf $DIST_DIR
mkdir $DIST_DIR

tar --create . | docker run --rm -i golang:$GOVERSION bash -c "$(cat build/docker_build.sh)" > ${DIST_DIR}/build.tar

cd $DIST_DIR

tar --extract --file=build.tar
rm -f build.tar

# Homebrew - binary must be called phraseapp, because the binary name inside
# the tar will be made available system wide
cp phraseapp_macosx_amd64 phraseapp
tar --create --mtime="@${SOURCE_DATE_EPOCH}" phraseapp | gzip -n > phraseapp_macosx_amd64.tar.gz
rm phraseapp

for name in phraseapp_linux_386 phraseapp_linux_amd64; do
  tar --create --mtime="@${SOURCE_DATE_EPOCH}" $name | gzip -n > ${name}.tar.gz
done

if ! which zip > /dev/null; then
  echo "zip not installed"
fi

zip phraseapp_windows_amd64.exe.zip phraseapp_windows_amd64.exe > /dev/null

echo "Last change: ${LAST_CHANGE}"
echo "Version: ${VERSION}"
echo "Brew hash: $(sha256sum phraseapp_macosx_amd64.tar.gz | cut -d ' ' -f 1)"
echo "Build output: $(pwd)"
