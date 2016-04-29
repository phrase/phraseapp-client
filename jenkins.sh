#!/bin/bash
set -e

export BUILD_DIR=$(dirname $0)
pushd $BUILD_DIR > /dev/null

export GOVERSION=${GOVERSION:-1.6.2}
export PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games
export REVISION=${GIT_COMMIT:-$(git rev-parse HEAD)}
export LIBRARY_REVISION=$(cat Godeps/Godeps.json | grep github.com/phrase/phraseapp-go -A 1 | tail -n 1 | cut -d '"' -f 4)
export PROJ_DIR=/go/src/github.com/phrase/phraseapp-client
export ORIGINAL_VERSION=$VERSION
export CURRENT_DATE=$(TZ=UTC date +"%Y-%m-%dT%H:%M:%SZ")

mkdir -p tmp
export DIR=$(mktemp -d ${PWD}/tmp/phraseap-client-XXXX)
if [[ "$KEEP_RELEASE" != "true" ]]; then
  trap "rm -Rf $DIR" EXIT
else
  echo "keeping release in $DIR"
fi

if [[ -z $LIBRARY_REVISION ]]; then
  echo "unable to get library revision"
  exit 1
fi

if [[ -z $VERSION ]]; then
  # try to fetch the most recent version and use <version>-dev
  VERSION=$(git log --pretty=format:'%d' | ruby -e 'puts STDIN.readlines.map { |l| l[/tag: ([\d\.]+)/, 1] }.compact.first')-dev
fi

# test and vet
docker run --rm -i -v "$PWD":$PROJ_DIR -w $PROJ_DIR golang:$GOVERSION bash <<EOF
set -xe
go test ./...
go vet ./...
EOF

echo "building version=${VERSION} revision=${REVISION} library_revision=${LIBRARY_REVISION}"

function build {
  goos=$1
  goarch=$2
  name=$3
  if [[ -z $name ]]; then
    echo "name must be present"
    exit 1
  fi
  echo "build os=${goos} arch=${goarch} name=$name"
  proj_dir=/go/src/github.com/phrase/phraseapp-client
  docker run --rm -i -e GOOS=$goos -e GOARCH=$goarch -v $DIR:/go/bin -v "$PWD":$proj_dir -w $proj_dir golang:$GOVERSION bash <<EOF
  set -e
  go build -o /go/bin/$name -ldflags "-X main.BUILT_AT=$CURRENT_DATE -X=main.REVISION=$REVISION -X=main.PHRASEAPP_CLIENT_VERSION=$VERSION -X=main.LIBRARY_REVISION=$LIBRARY_REVISION" .
EOF
}

build linux   amd64   phraseapp_linux_amd64
build linux   386     phraseapp_linux_386
build darwin  amd64   phraseapp_macosx_amd64
build windows amd64   phraseapp_windows_amd64.exe

pushd $DIR > /dev/null

# Homebrew - binary must be called phraseapp, because the binary name inside
# the tar will be made available system wide
cp phraseapp_macosx_amd64 phraseapp
tar cfz phraseapp_macosx_amd64.tar.gz phraseapp
rm phraseapp

for name in phraseapp_linux_386 phraseapp_linux_amd64; do
  tar cfz ${name}.tar.gz $name
done

zip phraseapp_windows_amd64.exe.zip phraseapp_windows_amd64.exe &> /dev/null
popd > /dev/null

if [[ -n $WORKSPACE ]]; then
  # probably running inside jenkins
  dst=s3://phraseapp-client-releases/${ORIGINAL_VERSION:-$REVISION}/
  aws s3 sync --delete --acl=public-read $DIR $dst
  shasum=$(sha256sum $DIR/phraseapp_macosx_amd64 | awk '{ print $1 }')
  if [[ -z $shasum ]]; then
    echo "unable to get shasum of phraseapp_macosx_amd64"
    exit 1
  fi
  aws s3 cp --acl=public-read --metadata SHA256=${shasum} $DIR/phraseapp_macosx_amd64  ${dst}phraseapp_macosx_amd64
fi
