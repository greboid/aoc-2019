#!/bin/bash

IMAGE=greboid/aoc-2019-02

docker image inspect $IMAGE >/dev/null 2>&1
if [ $? -ne 0 ]
then
    echo "Building docker image..."
    cd docker
    docker build . -t $IMAGE
    cd ..
fi

docker run --rm -it -v "$(pwd)":/app $IMAGE /entrypoint.sh $@