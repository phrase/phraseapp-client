#!/bin/bash

set -eo pipefail

# validate that master is checked out and head points to origin/master
BRANCH=$(git rev-parse --abbrev-ref HEAD)
if [ $BRANCH != 'master' ]; then
    echo "Not on master branch. Aborting"
    exit 1
fi

 HEADHASH=$(git rev-parse HEAD)
 UPSTREAMHASH=$(git rev-parse master@{upstream})

 if [ "$HEADHASH" != "$UPSTREAMHASH" ]; then
   echo "Not up to date with origin/master. Aborting"
   exit 1
 fi

# Validate current tag against version
current_tag=$(git describe --tags --abbrev=0)
client_version=$(cat .version)

if [ $current_tag != $client_version ];then
    echo "client version '${client_version}' and tag '${current_tag}' do not match"
    exit 1
fi

# Build client
./build/build.sh
./build/innosetup/create_installer.sh

# Create release
function create_release_data()
{
  cat <<EOF
{
  "tag_name": "${current_tag}",
  "name": "${current_tag}",
  "draft": true,
  "prerelease": false
}
EOF
}

echo "Create release $client_version"
api_url="https://api.github.com/repos/phrase/phraseapp-client/releases?access_token=${GITHUB_TOKEN}"
response="$(curl --data "$(create_release_data)" ${api_url})"
release_id=$(echo $response | python -c "import sys, json; print json.load(sys.stdin)['id']")

if [ -z "$release_id" ]
then
      echo "Failed to create GitHub release"
      echo $response
      exit 1
else
      echo "New release created created with id: ${release_id}"
fi

# Upload artifacts
DIST_DIR="./dist"
for file in "$DIST_DIR"/*; do
    echo "Uploading ${file}"
    asset="https://uploads.github.com/repos/phrase/phraseapp-client/releases/${release_id}/assets?name=$(basename "$file")&access_token=${GITHUB_TOKEN}"
    curl --data-binary @"$file" -H "Content-Type: application/octet-stream" $asset > /dev/null
done

echo "Release successful"
