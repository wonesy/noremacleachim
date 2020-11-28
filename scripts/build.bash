#!/bin/bash

pushd ../deployment
docker build -t wonesy/noremacleachim:latest . || exit 0

if [[ $1 == "publish" ]]; then
    docker push wonesy/noremacleachim:latest
fi

popd
