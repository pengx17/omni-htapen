#!/bin/bash

TAG="latest"

if [[ -n $1 ]]; then
    TAG=$1
fi

rm -rf dist
echo ">>> Building omni-htapen:build"

docker build -t pengxiao/omni-htapen:build -f ./dockerfiles/build.Dockerfile .
docker container create --name omni-htapen-build-extract pengxiao/omni-htapen:build
docker container cp omni-htapen-build-extract:./dist ./dist
docker container rm omni-htapen-build-extract

echo ">>> Building omni-htapen:$TAG"
docker build --no-cache -t pengxiao/omni-htapen:$TAG -f ./dockerfiles/Dockerfile .
docker push pengxiao/omni-htapen:$TAG
rm -rf dist
