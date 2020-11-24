#!/bin/bash

pushd ..
docker build -t wonesy/noremacleachim:latestgit . || exit 0

if [[ $1 == "publish" ]]; then
    docker push wonesy/noremacleachim:latest
fi

popd
