#!/bin/bash

export NAME=aws-secret-manager-utility
export VERSION=v1.0

if [ $1 ]
then
    export VERSION=$1
fi

docker build --tag aws-secret-manager-utility .
docker run --env GITHUB_TOKEN --env NAME --env VERSION aws-secret-manager-utility