#!/bin/bash
set -e

SRC_DIR=$GOPATH/src/github.com/phrase/phraseapp-client
mkdir -p $SRC_DIR
cd $SRC_DIR
tar --extract
bash ./build/go_test.sh
