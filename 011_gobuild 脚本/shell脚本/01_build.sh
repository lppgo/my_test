#!/usr/bin/env bash

function pullCode() {
    git pull --rebase
}


function build() {
    export GO111MODULE=on
    flags="-X main.buildstamp=`date '+%Y-%m-%d/%H:%M:%S/%p'` -X 'main.goversion=$(go version)' -X main.githash=`git rev-parse --short HEAD`"
    GOOS=linux go build -ldflags "$flags" -o cmd/mountain/mountain cmd/mountain/main.go
}


pullCode

if [ $? -ne 0 ]; then
    echo "pull failed"
else
    echo "pull succeed"
    build
