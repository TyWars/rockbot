#!/bin/bash

set -e

tag=$1
if [ -z "$tag" ]
then
    tag=latest
fi

. build.sh

docker build -t sonas-bundle .
docker tag sonas-bundle atiwari/sonas-bundle:$tag
docker push atiwari/sonas-bundle:$tag
