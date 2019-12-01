#!/bin/bash

set -e

cd /app

if [ -d "$1" ]; then
    if [ ! -f "${GOPATH}/bin/$1" ] || [ "${GOPATH}/bin/$1" -nt "$1/main.go" ]; then
        export HOME="/tmp"
        go get -d -v ./... >/dev/null
        go install ./... >/dev/null
    fi
    time "${GOPATH}/bin/$1"
else
    echo 'Day does not exist.'
fi