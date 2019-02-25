#!/bin/bash

TAG="latest"

if [[ -n $1 ]]; then
    TAG=$1
fi

rm -rf dist
echo ">>> Building omini-htapen:build"

docker build -t pengxiao/omini-htapen:build -f ./dockerfiles/build.Dockerfile .
docker container create --name omini-htapen-build-extract pengxiao/omini-htapen:build
docker container cp omini-htapen-build-extract:./server ./
docker container rm omini-htapen-build-extract

echo ">>> Building omini-htapen:$TAG"
docker build --no-cache -t pengxiao/omini-htapen:$TAG -f ./dockerfiles/Dockerfile .
rm -rf dist
